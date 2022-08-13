### Systems programming with Go
Systems programming is a special area of programming on Unix machines (and not limited to unix).Most commands that have to do with system
administration tasks, such as disk formatting, network interface configuration, module loading, and kernel performance tracking, are implemented using the techniques of systems programming. Additionally, the /etc directory, which can be found on all Unix systems, contains plain text files that deal with the configuration of a Unix machine and its services and are also manipulated using systems software.

We can group the various areas of systems software and related system calls in the following sets:
   * File I/O: This area deals with file reading and writing operations, which is the most important task of an operating system. File input and output must be fast and efficient, and above all, reliable.
   * Advanced file I/O: Apart from the basic input and output system calls, there are also more advanced ways to read or write to a file including asynchronous I/O and non-blocking I/O.
   * System files and configuration: This group of system software includes functions that allow you to handle system files, such as /etc/passwd , and get system specific information, such as system time and DNS configuration.
   * Files and directories: This cluster includes functions and system calls that allowthe programmer to create and delete directories and get information such as the owner and the permissions of a file or a directory.
   * Process control: This group of software allows you to create and interact with Unix processes.
   * Threads: When a process has multiple threads, it can perform multiple tasks. However, threads must be created, terminated, and synchronized, which is the purpose of this collection of functions and system calls.
   * Server processes: This set includes techniques that allow you to develop server processes, which are processes that get executed in the background without the need for an active terminal.
   *Interprocess communication: This set of functions allows processes that run on the same Unix machine to communicate with each other using features such as pipes, FIFOs, message queues, semaphores, and shared memory.
   * Signal processing: Signals offer processes a way of handling asynchronous events, which can be very handy.
   * Network programming: This is the art of developing applications that work over computer networks with the help of TCP/IP and is not systems programming per se. However, most TCP/IP servers and clients are dealing with system resources, users, files, and directories. So, most of the time, you cannot create network applications without doing some kind of systems programming.