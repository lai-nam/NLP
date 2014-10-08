function POST() {
    var category = document.getElementById("nlp_category").value;
    var content = document.getElementById("nlp_content").value;

    $.ajax({
        url: "/api/documents/",
        dataType: "application/json",
        data: JSON.stringify({
            category: category,
            content: content
        }),
        type: "POST",
        complete: function(err, data){
        	alert("ok");
        }
    });
}
