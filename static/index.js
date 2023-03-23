
let contentHolder = document.getElementById("content-holder");

function deletePost(postEl) {
    if (contentHolder.children.length <= 4) {
        load(document.getElementById("content-holder").lastChild.dataset.time)
    }
    postEl.remove();
}

function archivePost(postEl) {
    if (contentHolder.children.length <= 4) {
        load(document.getElementById("content-holder").lastChild.dataset.time)
    }
    postEl.remove();
}

function createPostEl(post) {
    let rootEl = document.createElement("div")
    rootEl.classList.add("swipe-root")
    rootEl.dataset.time = post.TimeProcessing

    let holderEl = document.createElement("div")
    holderEl.classList.add("box")

    let startX, startY, currentX, currentY, distX, distY;

    rootEl.addEventListener("touchstart", (event) => {
        const touch = event.touches[0];
        startX = touch.clientX;
        startY = touch.clientY;
    });

    rootEl.addEventListener("touchmove", (event) => {
        event.preventDefault();
        currentX = event.touches[0].clientX;
        currentY = event.touches[0].clientY;
        distX = currentX - startX;
        distY = currentY - startY;
        holderEl.style.transform = `translateX(${distX}px)`;
    });

    rootEl.addEventListener("touchend", (event) => {
        if (Math.abs(distX) > Math.abs(distY)) {
            console.log(distX);
            if (distX > holderEl.offsetWidth / 2) {
                deletePost(rootEl)
                return
            }
            if (distX < -(holderEl.offsetWidth / 2)){
                archivePost(rootEl)
                return
            }
        }
        holderEl.style.transform = `translateX(0px)`;
    });

    let laterBtn = document.createElement("button")
    laterBtn.innerText = "<"
    laterBtn.classList.add("swipe-btn")
    laterBtn.classList.add("flex-child")
    laterBtn.addEventListener("click", () => {
        archivePost(rootEl)
    })
    let deleteBtn = document.createElement("button")
    deleteBtn.innerText = ">"
    deleteBtn.classList.add("swipe-btn")
    deleteBtn.classList.add("flex-child")
    deleteBtn.addEventListener("click", () => {
        deletePost(rootEl)
    })

    let contentEl = document.createElement("div")
    contentEl.classList.add("flex-child")
    let titleEl = document.createElement("h3")
    titleEl.innerText = post.Title
    let textEl = document.createElement("div")
    textEl.innerText = post.Text
    let linkEl = document.createElement("a")
    linkEl.href = post.Url
    linkEl.innerText = "view"
    linkEl.target = "_blank";

    contentEl.appendChild(titleEl)
    contentEl.appendChild(textEl)
    contentEl.appendChild(linkEl)

    holderEl.appendChild(laterBtn)
    holderEl.appendChild(contentEl)
    holderEl.appendChild(deleteBtn)

    rootEl.appendChild(holderEl)

    return rootEl
}

function load(starttime) {
    fetch(`/load/${starttime}`).then(res => res.json()).then(data => {
        if(data.status != "success") {
            console.error("server returned error");
            return
        }
        data.data.forEach(post => {
            console.log(post);
            contentHolder.appendChild(createPostEl(post))
        });
    }).catch(err => console.error(err))
}

load(0)

document.getElementById("content-load-btn").addEventListener("click", () => {
    load(document.getElementById("content-holder").lastChild.dataset.time)
})
