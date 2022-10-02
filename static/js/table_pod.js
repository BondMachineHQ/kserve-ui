var table = $('#example').DataTable({
    serverSide: false,
    lengthChange: false,
    ajax: {
        url: "/list_pods",
        dataSrc: 'items'
    },
    columns: [
      { data: 'spec.predictor.model.modelFormat.name'},
      { data: 'spec.predictor.model.protocolVersion',
        defaultContent: "<i>Not available</i>"
      },
      { data: 'metadata.name'},
      { data: 'status.url',
        defaultContent: "<i>Not available yet</i>",
        render: function ( data, type, row) {
          if (typeof data !== 'undefined') {
            return data.replace(/^http?:\/\//, '');
          }
          return "<i>Not available yet</i>";
        }
      },
      { data: 'spec.predictor.model.storageUri',
        render: function ( data, type, row) {
          return '<a href="'+data+'">'+data+'</a>';
        }
      }
    ],
    columnDefs: [
            {
                targets: 5,
                data: null,
                defaultContent: '<button type="delete" class="btn waves-effect waves-light" >Delete<i class="material-icons right">delete</i></button>',
            },
     ],
});

$('#refresh_button').click(function refreshData() {
     table.ajax.reload();
})

$('#example tbody').on('click', 'button', function () {
    var data = table.row($(this).parents('tr')).data();
    //alert("Deleting " + data.metadata.name);
    $.ajax({
  
        // Our sample url to make request 
        url: "/delete_isvc",
        type: "POST",
        dataType:"json",
        contentType: "application/json",
        data: JSON.stringify({isvcname: data.metadata.name}),
        success: function (data) {
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
            table.ajax.reload();
         }, 5000)
    );
    alert(data.metadata.name + "has been deleted!")
});