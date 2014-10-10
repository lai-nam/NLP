function Word_POST() {
    var word = document.getElementById("nlp_word").value;
    var tags = document.getElementById("nlp_tags");

    var tag = tags.options[tags.selectedIndex].text;
   
    $.ajax({
        url: "/api/words/",
        dataType: "application/json",
        data: JSON.stringify({
            word: word,
            tag: tag
        }),
        type: "POST",
        complete: function(err, data) {
            alert("ok");
        }
    });
}
