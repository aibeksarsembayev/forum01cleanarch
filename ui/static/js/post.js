// post creation ...
function createPostJSON() {
  const title = document.getElementById("postTitle");
  const content = document.getElementById("postContent");
  const category = document.getElementById("category");

  // Crearing a XHR object
  let xhr = new XMLHttpRequest();
  let url = "/post/create";

  // open connection
  xhr.open("POST", url, true);

  // Set the request header i.e. which type of of content you are sending
  xhr.setRequestHeader("Content-Type", "application/json");

  // Create a state change callback
  xhr.onreadystatechange = function () {
   if (xhr.readyState === 4 && xhr.status === 200) {

    var jsonResponse = JSON.parse(xhr.responseText);

      // Print received data from server
      // result.innerHTML = this.responseText;

      window.location.replace("/post/" + jsonResponse.post_id)
    } 
  // need to add response in case of fail
  };
  // Converting JSON data to string
  var data = JSON.stringify({ "title": title.value, "content": content.value, "category_name": category.value});

  // Sending data with the request
  xhr.send(data)
}