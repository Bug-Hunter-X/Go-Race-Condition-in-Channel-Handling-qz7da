# Go Race Condition Example

This repository demonstrates a race condition in a Go program that uses goroutines and channels.  The program appears to work correctly on the surface but will exhibit unpredictable behavior due to the unhandled race condition.  The solution file shows how to correctly synchronize access to the channel using wait groups.