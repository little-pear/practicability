// ?获取展示文本的span 元素
const textEl = document.querySelector("#text");
const texts = JSON.parse(textEl.getAttribute("data-text"));

// ?获取数组第一行文本
let index = 0;
// ?每行文本第几个字符
let charIndex = 0;
// ?间隔时间
let delta = 500;
// ?记录动画开始执行的时间
let start = null;
let isDeleting = false;

function type(time) {
    window.requestAnimationFrame(type);
    if (!start) start = time;
    let progress = time - start;

    if (progress > delta) {
        let text = texts[index];
        if (!isDeleting) {
            textEl.innerHTML = text.slice(0, ++charIndex);
            delta = 500 - Math.random() * 400;
        } else {
            textEl.innerHTML = text.slice(0, charIndex--);
        }

        start = time;

        if (charIndex === text.length) {
            isDeleting = true;
            delta = 200;
            start = time + 1200;
        }

        if (charIndex < 0) {
            isDeleting = false;
            start = time + 200;
            index = ++index % texts.length;
        }
    }
}

window.requestAnimationFrame(type);
