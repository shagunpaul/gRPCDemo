A simple project to send messages between client to server via gRPC

gRPC (Remote Procedure Call)- framework enables efficient communication where client can directly call functions from client to server

Types of Calls that can be made via gRPC

1.Unary
client sends a single request to the server and receives a single response.

2.Server Streaming
the client sends a single request to the server, and the server responds with a stream of messages.

1.Client Streaming
the client sends a stream of messages to the server, and the server responds with a single response.

1.Bi Directional Streaming
both the client and server send a stream of messages asynchronously.