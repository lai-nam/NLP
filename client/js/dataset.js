document.addEventListener("DOMContentLoaded", function(event) {

    var el = document.getElementsByClassName('save');
    if (el) {
        el[0].addEventListener('click', function() {
            var doc = document.getElementById('document');
            var cl = document.getElementById('cl');
            if (doc.value) {
                console.log("doc.value: ", doc.value);
                $.ajax({
                    url: "/v1/training/",
                    dataType: "json",
                    type: "post",
                    contentType: "application/json",
                    data: JSON.stringify({
                        "document": doc.value,
                        "class": cl.value
                    }),
                    success: function(err, data) {
                        console.log("Err: ", err, data);
                    }
                })
            } else {
                console.log("Vui lòng nhập dữ liệu!")
            }
            // alert("TODO: Doc: " + doc.value + " Polarity: " + pol.value);
            // TODO: 
            // Gửi document và trọng polarity để lưu lại.
            // doc.value, pol.value
        }, false);
    }
});

function polarityChange() {
    var cl = document.getElementById('cl');
    if (cl.value == -1) {
        $(".icon-emotion").addClass('neg');
    } else {
        $(".icon-emotion").removeClass('neg');
    }
};
