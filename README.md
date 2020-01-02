# Demo of Potential Non-Greedy Bug in `peg`

Just a working example of what looks like a diversion from what I'm reading is a core PEG behavior: that they are greedy by default. Is this by design?

```
 element ?
              The  element  is  optional.  If present on the input, it is con-
              sumed and the match succeeds.  If not present on the  input,  no
              text is consumed and the match succeeds anyway.
```

<https://www.piumarta.com/software/peg/peg.1.html>
