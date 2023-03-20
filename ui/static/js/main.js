// Navigation
// Mobile Menu Toggle
// const hamburger = document.querySelector('.hamburger');
// const navigation = document.querySelector('.navigation');

// hamburger.addEventListener('click', function () {
//     navigation.classList.toggle('open');
// });

const toggleMenu = document.querySelector(".navigation button");
const menu = document.querySelector(".navigation ul");

toggleMenu.addEventListener("click", function () {
    const open = JSON.parse(toggleMenu.getAttribute("aria-expanded"));
    toggleMenu.setAttribute("aria-expanded", !open);
    menu.hidden = !menu.hidden;

    if (!menu.hidden) {
        menu.querySelector('a').focus();
    }
});


// Active page
const navLinks = document.querySelectorAll("nav ul li a")
for (var i = 0; i < navLinks.length; i++) {
    var link = navLinks[i];
    if (link.getAttribute("href") == window.location.pathname) {
        link.classList.add("active");
        break;
    }
}
