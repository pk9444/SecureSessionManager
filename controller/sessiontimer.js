var timer = document.getElementById("timer")
var start = 0;

setInterval(updateTime, 1000);

function updateTime(){
    start--;
    if(start < 1){
        window.Location("components/login.html");
    }else{
        timer.innerText = "Session Expires in : " + formatTimer(start)
    }
}
window.addEventListener("mousemove", reset)

function reset(){
    start=10;
}
function formatTimer(timeInSec){
    var min = Math.floor(timeInSec / 60);
    var sec = timeInSec % 60;
    if ( min < 10 ) { min = "0" + min; }
	if ( sec < 10 ) { sec = "0" + sec; }
	return min + ":" + sec;

}

