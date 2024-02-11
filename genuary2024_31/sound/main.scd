(
var cleanup = {
  "cleanup".postln;
  ServerTree.removeAll;
  ServerBoot.removeAll;
  ServerQuit.removeAll;
  CmdPeriod.run;
};

var makeBusses = {
  "makeBusses".postln;
  ~out = 0;
  s.newBusAllocators;
  ~reverbBus = Bus.audio(s, 2);
  ~delayBus = Bus.audio(s, 2);
  ~masterBus = Bus.audio(s, 2);
};

var makeSynthDefs = {
  "makeSynthDefs".postln;
  thisProcess.interpreter.executeFile("/home/keith/Code/supercollider/lib/synthdefs.scd");
  thisProcess.interpreter.executeFile("/home/keith/Code/supercollider/lib/effects.scd");
  ~keyboardSynth = \pulse;
};

var makeEffects = {
  "makeEffects".postln;
  ~master = Synth(\masterAmp, [in: ~masterBus, out: ~out, amp: 0.5]);
  ~delay = Synth(\delay, [in: ~delayBus, out: ~masterBus, delay: 0.15, decay: 2, drive: 0.3]);
  ~reverb = Synth(\reverb, [in: ~reverbBus, out: ~delayBus, room: 0.7]);
};

var postDone = {
  "=================================".postln;
  "= READY!!!!!".postln;
  "=================================".postln;
};

s.waitForBoot({
  cleanup.();
  makeBusses.();
  makeSynthDefs.();
  s.sync;
  thisProcess.interpreter.executeFile("midi.scd");
  ~setupMidi.(false);
  ~addMidiEvents.();
  ServerTree.add(makeEffects);
  ServerTree.add(~addMidiEvents);
  ServerTree.add(postDone);

});
)

// can change this easily.
~keyboardSynth = \kick;

(
Pbindef(\main,
  \instrument, \bass,
  \dur, Prand([0.5, 0.5], inf),
  \octave, 4,
  \degree, Prand((0..8), inf),
  \amp, 0.5,
  \out, ~reverbBus,
);
)
Pbindef(\main).play
Pbindef(\main).stop

(
Pbindef(\pad,
  \instrument, \pad,
  \dur, Prand([2.0, 4.0], inf),
  \octave, 5,
  \degree, Prand((0..8), inf),
  \amp, 0.02,
  \out, ~reverbBus,
);
)
Pbindef(\pad).play
Pbindef(\pad).stop

s.record
s.stopRecording

~reverb.set(\room, 0.7)
~delay.set(\delay, 0)
