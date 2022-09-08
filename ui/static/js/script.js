// user signup ...
function signupJSON(){
  let result = document.querySelector('.result')

  const username = document.getElementById("username");
const email = document.getElementById("emailAddress");
const password = document.getElementById("userPassword");

// Crearing a XHR object
let xhr = new XMLHttpRequest();
let url = "/signup";

// open connection
xhr.open("POST", url, true);

// Set the request header i.e. which type of of content you are sending
xhr.setRequestHeader("Content-Type", "application/json");

// Create a state change callback
xhr.onreadystatechange = function () {
  if (xhr.readyState === 4 && xhr.status === 201) {

    // Print received data from server
    // result.innerHTML = this.responseText;
    window.location.replace("/signin")

  } 
  // need to add response in case of fail
};

// Converting JSON data to string
var data = JSON.stringify({ "username": username.value, "email": email.value, "password": password.value});

// Sending data with the request
xhr.send(data)
}

// user signin ...
function signinJSON() {
const email = document.getElementById("emailAddress");
const password = document.getElementById("userPassword");

// Crearing a XHR object
let xhr = new XMLHttpRequest();
let url = "/signin";

// open connection
xhr.open("POST", url, true);

// Set the request header i.e. which type of of content you are sending
xhr.setRequestHeader("Content-Type", "application/json");

// Create a state change callback
xhr.onreadystatechange = function () {
  if (xhr.readyState === 4 && xhr.status === 200) {

    // Print received data from server
    // result.innerHTML = this.responseText;
    window.location.replace("/")

  } 
  // need to add response in case of fail
};

// Converting JSON data to string
var data = JSON.stringify({ "email": email.value, "password": password.value});

// Sending data with the request
xhr.send(data)
  
}

