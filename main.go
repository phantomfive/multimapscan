package main
import ("fmt";"sync/atomic";"sync";"time";"os/exec";"os")


var runningThreads int32 = 0
var MAX_THREADS int32 = 300

func main() {

	for i:=1; i<len(os.Args); i++ {
		time.Sleep(10 * time.Millisecond)
		for ;runningThreads>MAX_THREADS; {
			time.Sleep(100 * time.Millisecond)
		}
		atomic.AddInt32(&runningThreads, 1)
		go probePort(os.Args[i])
	}
	for ; runningThreads > 0; {
	}
}

func probePort(ip string) {

	fmt.Fprint(os.Stderr,".")
	cmd :=exec.Command("nmap", "-Pn", "-T5",  ip)
	out, e := cmd.Output()
	if e==nil {
		fmt.Fprintf(os.Stderr, "mapped IP: %s", ip)
		output(string(out))
	}else {
		fmt.Fprintf(os.Stderr, "err: %s    %s", out, e)
	}
	
	atomic.AddInt32(&runningThreads, -1)
}

var outputMutex sync.Mutex
func output(s string) {
	outputMutex.Lock()

	fmt.Println(s)
	
	outputMutex.Unlock()
}




