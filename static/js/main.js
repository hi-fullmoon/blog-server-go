$(function () {
    // loading start
    NProgress.start();

    // set scrollbar
    $('body').niceScroll({cursorwidth: 8, cursorcolor: '#778886'});

    // search
    $('#search-input').on('click', function () {
        $('.js-search-panel').show();
    });

    // search input closed
    $('.js-close').on('click', function () {
        $('.js-search-panel').hide();
        $('.js-search-input').val('');
        $('.js-search-result').empty();
    });

    $('.js-search-input').on('input', debounce(function () {
        getArticles();
    }, 500));

    // get articles
    function getArticles() {
        var value = $('.js-search-input').val();
        if (value === '') {
            $('.js-search-result').empty();
            return
        }

        $.ajax({
            type: 'GET',
            url: '/api/user/articles?title=' + value,
            dataType: 'json',
            success: function (res) {
                if (res.code === 0) {
                    $('.js-search-result').empty();
                    var nodeArr = [];
                    res.data.forEach(item => {
                        nodeArr.push(
                            '<li class="search-result__item"><a href="/articles/' + item.id + '">' + item.title + '</a></li>'
                        )
                    });

                    if (nodeArr.length === 0) {
                        nodeArr.push(
                            '<li style="padding: 0 20px;" class="search-result__item">抱歉，木有查到你想要的文章···</li>'
                        )
                    }

                    $('.js-search-result').append(nodeArr.join(''));
                }
            }
        });
    }

    // set search result container scrollbar
    $('.js-search-result').niceScroll({
        cursorwidth: 8,
        background: '#eee',
        cursorcolor: '#78887d',
        railpadding: {top: 0, right: 0, left: 0, bottom: 0},
        autohidemode: false,
        cursorborderradius: 0,
        cursorborder: 0
    });

    // loading done
    NProgress.done();
});

/**
 * 防抖函数
 * @param method 事件触发的操作
 * @param delay 多少毫秒内连续触发事件，不会执行
 * @returns {Function}
 */
function debounce(method, delay) {
    var timer = null;
    return function () {
        var self = this;
        var args = arguments;

        timer && clearTimeout(timer);

        timer = setTimeout(function () {
            method.apply(self, args);
        }, delay);
    }
}