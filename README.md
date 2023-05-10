# Multithreaded TCP Server
 - TCP is the most reliable way for two machines to communicate to each other over the network.

## What is a TCP Server?

  - It is a simple process that runs on a machine and 'listens' to a port
  - Any machine that needs/wants to talk to server has to connect over the port and establish the connection.

 ## Steps to create a single threaded TCP server
    
   ### Start listening to the port
    - When process starts, pick a port and start listening to it.
   ### Wait for a client to connect
   
    - Invoke the **'ACCEPT'** system call and wait for a client to connect.
    - **'ACCEPT'** is a blocking call and server wouldnot proceed until some client connects.

   ### Read the request and send the response

    - Once the connection is established, invoke the **'READ'** system call to read the request.
    - Invoke the **'WRITE'** system call to send the response.
    - Both **'READ'** and **'WRITE'** are blocking calls.
    - Close the connection.

   ### Do this over and over
   - continiously waiting for client to connect, reading the request, writing the response and closing the connecion.

   ## Serving Multiple request concurrently
    -  once client connects, fork a thread to process and respond
    -  Ler main thread come back to **'ACCEPT'** as quickly as possible.

    ## Improvements: 
        * Limiting the number of threads.
        * Added thread pool to save on thread creation time
        * Connection timeout
        * TCP backlog queue configuration (OS LEVEL).