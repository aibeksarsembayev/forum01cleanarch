function createCommentJSON() {
  const result = document.querySelector('.resultComment')

  const comment = document.getElementById("comment_content")
  const post = document.getElementById("postID")

  // const comment = document.createComment("My comments")
  // document.body.appendChild(comment)

   // Crearing a XHR object
   let xhr = new XMLHttpRequest();
   let url = "/post/comment";

    // open connection
  xhr.open("POST", url, true);

  // Set the request header i.e. which type of of content you are sending
  xhr.setRequestHeader("Content-Type", "application/json");

  // Create a state change callback
  xhr.onreadystatechange = function () {
    if (xhr.readyState === 4 && xhr.status === 201) {

      // var jsonResponse = JSON.parse(xhr.responseText);

      // Print received data from server
      // result.innerHTML = this.responseText;

      window.location.replace("/post/" + post.value)
    } else {   // need to add response in case of fail
      var jsonResponse = JSON.parse(xhr.responseText);
      result.innerHTML = jsonResponse.message
    }
  };

  // Converting JSON data to string
  var data = JSON.stringify({ "comment_content": comment.value, "post_id": post.value});

  // Sending data with the request
  xhr.send(data)
}