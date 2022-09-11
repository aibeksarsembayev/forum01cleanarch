// post vote ...
function votePostLike() {
  let result = document.querySelector('.resultPostLike')

  const post_id = document.getElementById("postIDLike");
  const vote = document.getElementById("voteLike");

  // Crearing a XHR object
  let xhr = new XMLHttpRequest();
  let url = "/post/vote/";

  // open connection
  xhr.open("POST", url, true);

  // Set the request header i.e. which type of of content you are sending
  xhr.setRequestHeader("Content-Type", "application/json");

  // Create a state change callback
  xhr.onreadystatechange = function () {
    if (xhr.readyState === 4 && xhr.status === 201) {

      // Print received data from server
      var jsonResponse = JSON.parse(xhr.responseText);
      result.innerHTML = jsonResponse.post.VoteLike
      // window.location.replace("/signin")
      console.log("option 1", data)
    } else { // need to add response in case of fail
      // if xhr.status === 401 {  }
      var jsonResponse = JSON.parse(xhr.responseText);
      result.innerHTML = jsonResponse.message
      console.log("option 2", data)

    }

  };

  // Converting JSON data to string
  var data = JSON.stringify({ "post_id": post_id.value, "post_vote_value": vote.value });

  // Sending data with the request
  xhr.send(data)
}