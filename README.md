# chatserver
A simple chat-server implemented using raw tcp sockets. This is a working prototype/POC for experimental purposes.

It reuses the same socket that was created during the clients registration throughout a single session hence providing real-time pushes from the server.

### How to run...?
Just do `go build` and launch the compiled binary.

### Usage & Test
You will need at least 2 TCP clients, REST clients will not work and they add a lot of header information that the server is not expecting.

1. Start the server. It will start up at port 8000.
2. Register the clients. Registration process is very trivial and explained below.
    1. Send a TCP request to the server with a single line of plain text(used for client name) as payload.
    2. The server reads a single line from the request and uses this string as a name for the new client.
    3. Server creates a new client with a passed name. Registration is now complete.
    4. After creating the client, the server does NOT close the socket. Instead, it starts listening over the same socket for messages.
3. Send message to other client(s).
    1. After registration, clients can exchange messages. A message has the following JSON format:
    ```
    {
      "Recipients": ["client1", "client2"...],
      "Content": "Hi, I'm a cow",
    }
    ```
    2. A JSON encoded single-line(uglified) string represents a single message, the server considers each line to an individual message.
    3. `Recipients` specify the recipients of this message as expected and `Content` is the message content.
    4. To send a message, just write a message string(in the above format) into the already existing socket(no new connection required). This can be throughout a single chat session.
