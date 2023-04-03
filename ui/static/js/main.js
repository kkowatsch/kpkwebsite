// Navigation

// Mobile Menu Toggle
const toggleMenu = document.querySelector(".navigation button");
const menu = document.querySelector(".navigation ul");

if (window.innerWidth < 800) {

    toggleMenu.addEventListener("click", function () {
        const open = JSON.parse(toggleMenu.getAttribute("aria-expanded"));
        toggleMenu.setAttribute("aria-expanded", !open);
        menu.hidden = !menu.hidden;

        if (!menu.hidden) {
            menu.querySelector('a').focus();
        }
    });
}


// Active page
const navLinks = document.querySelectorAll("nav ul li a")
for (var i = 0; i < navLinks.length; i++) {
    var link = navLinks[i];
    if (link.getAttribute("href") == window.location.pathname) {
        link.classList.add("active");

        var active = document.querySelector(".active");
        var activeParent = active.parentElement.parentElement.previousElementSibling;

        activeParent.classList.add("active")
        break;
    }
}
