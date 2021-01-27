var request = new Request("http://127.0.0.1:8080/bilibili/register", {
    body:JSON.stringify({username:fm_username}),
    headers: new Headers({
    "Content-Type": "application/x-www-form-urlencoded"
    }),
    method : "POST",
    mode: "cors",
    redirect : "follow"
});
fetch(request)
.then((response) => {console.log(response);})
.catch((error)=>{console.log(error);});