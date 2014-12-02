$(".editbox").hide();

function DETable(id) {
    this.version = 'v1.0';
    this.id = id;
};
DETable.prototype.call = function(url, callback) {
    $.ajax({
        url: url,
        type: 'GET',
        success: function(data) {
            if (callback && typeof(callback) === "function") {
                callback(null, JSON.parse(data));
            }
        },
        error: function(e) {
            if (callback && typeof(callback) === "function") {
                callback("ERROR: ajax isn't successful", null);
            }
        }
    });
};
DETable.prototype.draw = function(data) {
    var self = this;
    this.data = data;

    this.table = $(this.id).DataTable({
        "data": data,
        "bPaginate": false,
        "bInfo": false,
        "columns": [{
            "data": "headword"
        }, {
            "data": "category"
        }, {
            "data": "subcategory"
        }, {
            "data": "defination",
            width: '25%'
        }, {
            "data": "eng"
        }, {
            "data": "tag"
        }, {
            "data": "wordtype"
        }]
    });
    console.log(this.id + ' tbody');
    $(this.id + ' tbody').on('click', 'tr', function() {
        if ($(this).hasClass('selected')) {
            $(this).removeClass('selected');
        } else {
            self.table.$('tr.selected').removeClass('selected');
            $(this).addClass('selected');
        }
    });
    this.fn = $(this.id).dataTable();
};
DETable.prototype.redraw = function(newData) {
    this.fn.fnClearTable();
    this.fn.fnAddData(newData);
    this.fn.fnDraw();
};
DETable.prototype.get = function() {
    return this.table.row('.selected').data();
};
DETable.prototype.index = function() {
    return this.table.row('.selected').index();
};
DETable.prototype.set = function(data, callback) {
    console.log("/v1/editor/" + data._id);
    $.ajax({
        url: "/v1/editor/" + data._id,
        type: "POST",
        contentType: 'application/json',
        data: JSON.stringify(data),
        dataType: 'json',
        success: function(data, textStatus, jqXHR) {
            if (callback && typeof(callback) === "function") {
                callback(1);
            }
        },
        error: function(jqXHR, textStatus, errorThrown) {
            if (callback && typeof(callback) === "function") {
                callback(0);
            }
        }
    });

};

var lv = new DETable('#table');
lv.call('/v1/editors/?per_page=100&page=1', function(err, data) {
    if (data) {
        lv.draw(data);
        var options = {
            currentPage: 1,
            totalPages: 416,
            onPageClicked: function(e, originalEvent, type, page) {
                lv.call('/v1/editors/?per_page=100&page=' + page, function(err, data) {
                    if (data) {
                        lv.redraw(data);
                    }
                });
            }
        };
        $('#panigation').bootstrapPaginator(options);
    }
});
$("#btnEdit").click(function() {
    var data = lv.get();
    if (data) {
        var content = $(".editbox .content");
        content.html("");
        _.each(data, function(v, k) {
            if (k != '_id') {
                var html = "<div class='row'>" +
                    "<label for='ip1' class='lbname col-sm-4'>" + k + "</label>" +
                    "<div class='col-sm-8'>" +
                    "<input key='" + k + "'class='form-control' id='ip1' placeholder='Text' value='" + v + "'>" +
                    "</div>" +
                    "</div>";
                content.append(html)
            }
        });
        $('.editbox .content input').on('keyup', function() {
            _t = $(this);
            var key = this.getAttribute('key');
            data[key] = _t.val();
        });
        $(".editbox").show();
        $("#btnSave").click(function() {
            JSON.stringify(data);

            lv.set(data, function(status) {
                if (status) {
                    $(".editbox").hide();
                } else {

                }
            });
        });
        $("#btnCancel").click(function() {
            $(".editbox").hide();
        });
    }
});
$("#btnDelete").click(function() {
    var data = lv.get();
    if (data) {

    }
});
