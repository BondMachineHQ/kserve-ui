$('#example').DataTable({
    serverSide: false,
    ajax: {
        url: "/list_pods",
        dataSrc: 'items'
    },
    columns: [
      { data: 'metadata.name'},
]
});