function fibonacci() {
    let num = parseInt(window.prompt("Enter a number", ""));
    if(!Number.isInteger(num)) {
        document.getElementById("output").innerHTML = "Please enter a valid number! (1 - ∞)"
        return;
    } else if (num === 1) {
        document.getElementById("output").innerHTML = "1"
        return;
    }
    let num1 = 0;
    let num2 = 1;
    let sum;
    let i = 0;
    let sequence = "";
    for (i = 0; i < num; i++) {
        sum = num1 + num2;
        if(i + 1 === num) {
            sequence = String(sequence) + String(sum)
        } else {
            sequence = String(sequence) + String(sum) + ", " 
        }
        
        num1 = num2;
        num2 = sum;
    }
    
    document.getElementById("output").innerHTML = String(sequence)
}

function getTime() {
    /*let today = new Date();
    document.getElementById("output").innerHTML = today.getHours() + ":" + today.getMinutes() + ":" + today.getSeconds()*/
    window.location.href = "http://localhost:9999";
}

function redBackground () {
    if(document.getElementById("output").style.backgroundColor == "red") {
        document.getElementById("output").style.backgroundColor = "rgb(5, 6, 45)";
    } else {
        document.getElementById("output").style.backgroundColor = "red"
    }
    
}

function redFont () {
    if(document.getElementById("output").style.color == "red") {
        document.getElementById("output").style.color = "white";
    } else {
        document.getElementById("output").style.color = "red"
    }
}
