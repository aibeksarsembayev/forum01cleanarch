function voteComment( vote, comment_id) {

  let resultl = document.querySelector('.' + 'likeComment' + comment_id)
  let resultd = document.querySelector('.' + 'dislikeComment' + comment_id)

  // Crearing a XHR object
  let xhr = new XMLHttpRequest();
  let url = "/post/comment/vote/";

  // open connection
  xhr.open("POST", url, true);

  // Set the request header i.e. which type of of content you are sending
  xhr.setRequestHeader("Content-Type", "application/json");

  // Create a state change callback
  xhr.onreadystatechange = function () {
    if (xhr.readyState === 4 && xhr.status === 201) {
      // Print received data from server
      var jsonResponse = JSON.parse(xhr.responseText);
      // console.log(jsonResponse)
      resultl.innerHTML = jsonResponse.comment_vote_like
      resultd.innerHTML = jsonResponse.comment_vote_dislike    
    } else if (xhr.readyState === 4 && xhr.status === 401) {
      window.location.replace("/signin")
    } else { // need to add response in case of fail  
    }
  };

  // Converting JSON data to string
  var data = JSON.stringify({ "comment_id": comment_id, "comment_vote_value": vote });

  // Sending data with the request
  xhr.send(data)
}