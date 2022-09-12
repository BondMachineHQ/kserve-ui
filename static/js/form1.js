window.addEventListener('submit', (event) => {

    event.preventDefault();

    const data = new FormData(event.target);

    const value = data.get('email');

    //location = "http://localhost:3000/list_pods?email=" + value

    $.ajax({
  
        // Our sample url to make request 
        url: "/create_pod",
        type: "POST",
        success: function (data) {
            var x = JSON.stringify(data);
            console.log(x);
        },

        // Error handling 
        error: function (error) {
            console.log(`Error ${error}`);
        }
    });

    console.log({value});
    console.log(value);
  }
)