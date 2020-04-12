import "../../Layout/index.css"
import * as $ from 'jquery'

$(".js-sign-out").click(function () {
    $.post("/signout", {
    }, function (response) {
        document.location = "/";
    }, "json")
        .fail(function(xhr) {
            error("Internal system error")
        });

    return false;
});


$(".user-list > div").click(function () {
    document.location = '/user/' + $(this).data('id')
});

$(".search").submit(function () {
    document.location = '/?name=' + $(this).find('input').val();
    return false;
});

$(".js-new-post-submit").click(function () {
    var text = $(".js-new-post-text").html()
    $.post("/feeds", {
        text: text
    }, "json")
        .done(function(response) {
            var feed = $("<div class='feed'>" + text + "</div>")
            $("js-feeds").prepend(feed)
        })
        .fail(function(xhr) {
            error("Internal system error")
        });
});

function loadFeeds() {
    var feed = $("<div class='feed'>" + text + "</div>")
    $.get("/feeds")

    $("js-feeds").prepend(feed)
}

$(function () {
    loadFeeds()
})
$(".feed")
