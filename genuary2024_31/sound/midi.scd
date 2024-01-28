(
~setupMidi = {
  arg reset = false;
  var midimix, mpkmini;
  if (~midiIsSetup != true || reset, {
    "MIDI: setupMidi... (wait)".postln;

    MIDIClient.init(verbose: false);
    MIDIIn.connectAll(false);
    MIDIdef.freeAll;
    midimix = MIDIClient.sources.detect({|e| e.device == "MIDI Mix"});
    if (midimix != nil, {
      "MIDI: found midimix!".postln;
      ~midimix = midimix.uid;
      ~midiMixOut = MIDIOut(0);
      ~midiMixOut.connect(~midimix);
    });

    mpkmini = MIDIClient.sources.detect({|e| e.device == "MPK mini 3"});
    if (mpkmini != nil, { 
      ~mpkmini = mpkmini.uid;
      "MIDI: found mpkmini!".postln;
    });
    ~midiIsSetup = true;
  });
};


~addMidiEvents = {
  var amp, notes, calc, mixBtns, toggleButton;
  "addEvents".postln;

  amp = 0.5;
  notes = [] ! 128;
  mixBtns = [] ! 28;

  toggleButton = {
    arg num;
    if (mixBtns[num] == 127, 
      { mixBtns[num] = 0; },
      { mixBtns[num] = 127; }
    );
    ~midiMixOut.noteOn(0, num, mixBtns[num]);
  };

  /////////////////////////////////////////////
  // knobs and sliders on mpkmini or midimix
  /////////////////////////////////////////////
  MIDIdef.cc(\controller, {
    arg val, num, chan, src;

    case
    {num==62} {
      calc = (val / 127).round(0.001);
      ~master.set(\amp, calc);
      postf("master amp: %\n", calc);
    }
    {num==16} {
      calc = val.linlin(0, 127, 0.1, 1);
      ~delay.set(\delay, calc);
      postf("delay:delay: %\n", calc);
    }
    {num==17} {
      calc = (val / 127 * 10).round(0.1);
      ~delay.set(\decay, calc);
      postf("delay:decay: %\n", calc);
    }
    {num==19} {
      calc = (val / 127).round(0.001);
      ~delay.set(\drive, calc);
      postf("delay:drive: %\n", calc);
    }
    {num==20} {
      calc = (val / 127).round(0.001);
      ~reverb.set(\room, calc);
      postf("reverb:room: %\n", calc);
    }
    {num==21} {
      calc = (val / 127).round(0.001);
      ~reverb.set(\damp, calc);
      postf("reverb:damp: %\n", calc);
    }
    {num==23} {
      calc = (val / 127).round(0.001);
      ~reverb.set(\mix, calc);
      postf("reverb:mix: %\n", calc);
    }
    {true} {};
  });

  /////////////////////////////////////////////
  // Pads
  /////////////////////////////////////////////
  MIDIdef.noteOn(\padon, {
    arg val, num, chan, src;
    case
    {num == 20} { 
      if (Pbindef(\main).isPlaying,
        { Pbindef(\main).stop; },
        { Pbindef(\main).play; }
      );
    }
  }, chan: 9, srcID: ~mpkmini);

  MIDIdef.noteOff(\padoff, {
    arg val, num, chan, src;
  }, chan: 9, srcID: ~mpkmini);


  /////////////////////////////////////////////
  // Keys
  /////////////////////////////////////////////
  MIDIdef.noteOn(\keyon, {
    arg val, num, chan, src;
    var note = Synth(~keyboardSynth, [freq: num.midicps, out: ~reverbBus]);
    notes[num] = note;
    NodeWatcher.register(note, true);
  }, chan: 0, srcID: ~mpkmini);

  MIDIdef.noteOff(\keyoff, {
    arg val, num, chan, src;
    var note = notes[num];
    if (note.isPlaying, {
      note.release;
      NodeWatcher.unregister(note);
    });

  }, chan: 0, srcID: ~mpkmini);

  /////////////////////////////////////////////
  // Midimix buttons
  /////////////////////////////////////////////
  MIDIdef.noteOn(\btnon, {
    arg val, num, chan, src;
    toggleButton.(num);
  }, srcID: ~midimix);

  MIDIdef.noteOff(\btnoff, {
    arg val, num, chan, src;
  }, srcID: ~midimix);
};
)
