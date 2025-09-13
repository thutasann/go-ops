# Channels

## Think of channels as doors 🚪

- Unbuffered channel = a door that can only be opened if both people (sender + receiver) are at the door at the same time.

- Buffered channel (size N) = a door with a small waiting room inside. Sender can drop things in the room until it’s full, then they must wait. Receiver can come later and pick them up.

## Blocking rules

- Send (ch <- v) blocks if:

  - Channel is unbuffered and no one is receiving yet.
  - Channel is buffered and the buffer is full.

- Receive (<-ch) blocks if:

  - Channel is unbuffered and no one is sending yet.
  - Channel is buffered and the buffer is empty.
