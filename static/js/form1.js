$(document).ready(function() {
    $('select').formSelect();
    // Old way
    // $('select').material_select();
});

window.addEventListener('submit', (event) => {

    event.preventDefault();

    const isvc_name = document.getElementById('isvc').value

    console.log(isvc_name)

    //location = "http://localhost:3000/list_pods?email=" + value

    $.ajax({
  
        // Our sample url to make request 
        url: "/create_pod",
        type: "POST",
        dataType:"json",
        contentType: "application/json",
        data: JSON.stringify({isvcname: isvc_name}),
        success: function (data) {
            console.log('${data}');
            M.toast({html: data.message});
        },

        // Error handling 
        error: function (error) {
            //console.log(`Error ${error}`);
            var x = JSON.parse(error.responseText);
            M.toast({html: x.message});
        }
    }).done(
        setTimeout(function(){
            window.location.reload();
         }, 5000)
    )

  }
)