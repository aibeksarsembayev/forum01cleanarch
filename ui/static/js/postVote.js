// post vote ...
function votePost(post_id, vote) {

  // console.log(post_id, vote)
  // const post_id = document.getElementById("postIDLike");
  // const vote = document.getElementById("voteLike");
  // const vote = input;


  let resultl = document.querySelector('.' + 'like' + post_id)
  let resultd = document.querySelector('.' + 'dislike' + post_id)

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
      // console.log(jsonResponse)
      resultl.innerHTML = jsonResponse.vote_like
      resultd.innerHTML = jsonResponse.vote_dislike
      // window.location.replace("/signin")

    } else if (xhr.readyState === 4 && xhr.status === 401) {
      window.location.replace("/signin")
    } else { // need to add response in case of fail
      // if xhr.status === 401 {  }     
      // var jsonResponse = JSON.parse(xhr.responseText);
      // result.innerHTML = jsonResponse.message
    }

  };

  // Converting JSON data to string
  var data = JSON.stringify({ "post_id": post_id, "post_vote_value": vote });

  // Sending data with the request
  xhr.send(data)
}

// function votePostDislike() {
//   let result = document.querySelector('.resultPostLike')
//   let resultd = document.querySelector('.resultPostDislike')

//   const post_id = document.getElementById("postIDDislike");
//   const vote = document.getElementById("voteDislike");

//   // Crearing a XHR object
//   let xhr = new XMLHttpRequest();
//   let url = "/post/vote/";

//   // open connection
//   xhr.open("POST", url, true);

//   // Set the request header i.e. which type of of content you are sending
//   xhr.setRequestHeader("Content-Type", "application/json");

//   // Create a state change callback
//   xhr.onreadystatechange = function () {
//     if (xhr.readyState === 4 && xhr.status === 201) {

//       // Print received data from server
//       var jsonResponse = JSON.parse(xhr.responseText);
//       console.log(jsonResponse)
//       result.innerHTML = jsonResponse.vote_like
//       resultd.innerHTML = jsonResponse.vote_dislike
//       // window.location.replace("/signin")

//     } else if (xhr.readyState === 4 && xhr.status === 401) {
//       window.location.replace("/signin")
//     } else { // need to add response in case of fail
//       // if xhr.status === 401 {  }     
//       // var jsonResponse = JSON.parse(xhr.responseText);
//       // result.innerHTML = jsonResponse.message
//     }

//   };

//   // Converting JSON data to string
//   var data = JSON.stringify({ "post_id": post_id.value, "post_vote_value": vote.value });

//   // Sending data with the request
//   xhr.send(data)
// }