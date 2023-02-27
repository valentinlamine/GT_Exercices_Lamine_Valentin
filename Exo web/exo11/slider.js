const slider = document.querySelector('.slider-width'); //select the slider
const item = document.querySelectorAll('.item'); //select all the items
const btnPrev = document.querySelector('.btn-prev'); //select the prev button
const btnNext = document.querySelector('.btn-next'); //select the next button
const itemWidth = 450; //the width of each item
const itemHeight = 400; //the height of each item

let itemDisplayed = 0; //the number of item displayed on the screen
let margin = 0; //the margin of the items
let sliderWidth = 0; //the width of the slider
let sliderPosition = 0; //the position of the slider

//function to calculate the number of item displayed on the screen
function calcItemDisplayed() {
    minimumMargin = 15; //the minimum margin of the items
    itemDisplayed = Math.floor((window.innerWidth - 100) / (itemWidth + minimumMargin)); //calculate the number of item displayed on the screen
    margin = Math.floor((((window.innerWidth - 100) - (itemDisplayed * itemWidth)) / (itemDisplayed))/2);  //calculate the margin of the items
    if (margin < minimumMargin) { //if the margin is too small
        itemDisplayed -= 1; //remove one item from the screen
        margin = Math.floor((((window.innerWidth - 100) - (itemDisplayed * itemWidth)) / (itemDisplayed))/2); //recalculate the margin
    }
    for (let i = 0; i < item.length; i++) { //apply the new margin to all the items
        item[i].style.marginLeft = margin + "px";
        item[i].style.marginRight = margin + "px";
    }
    //if all items are displayed on the screen mask the prev and next button
    if (itemDisplayed === item.length) {
        btnPrev.style.display = "none";
        btnNext.style.display = "none";
    } else {
        btnPrev.style.display = "block";
        btnNext.style.display = "block";
    }
}

//function to calculate the width of the slider
function calcSliderWidth() {
    sliderWidth = (itemWidth + (margin * 2)) * item.length; //calculate the width of the slider
    slider.style.width = sliderWidth + "px"; //apply the new width
}

//function to calculate the position of the slider
function calcSliderPosition() {
    sliderPosition = (itemWidth + (margin * 2)) * itemDisplayed; //calculate the position of the slider
}

//function to move the slider to the left
function prev() {
    if (sliderPosition < 0) { //if the slider is not at the beginning
        sliderPosition += (itemWidth + (margin * 2)); //move the slider to the left
        slider.style.transform = "translateX(" + sliderPosition + "px)"; //apply the new position
    } else { //if the slider is at the beginning
        sliderPosition = -((itemWidth + (margin * 2)) * (item.length - itemDisplayed)); //move the slider to the end
        slider.style.transform = "translateX(" + sliderPosition + "px)"; //apply the new position
    }
}

//function to move the slider to the right
function next() {
    if (sliderPosition > -((itemWidth + (margin * 2)) * (item.length - itemDisplayed))) { //if the slider is not at the end
        sliderPosition -= (itemWidth + (margin * 2)); //move the slider to the right
        slider.style.transform = "translateX(" + sliderPosition + "px)"; //apply the new position  
    } else { //if the slider is at the end
        sliderPosition = 0; //move the slider to the beginning 
        slider.style.transform = "translateX(" + sliderPosition + "px)"; //apply the new position
    }
}

//function to print all stats in the console
function printStats() {
    console.log("--------------------")
    console.log("(window.innerWidth - 100px) :", (window.innerWidth - 100));
    console.log("itemDisplayed :", itemDisplayed);
    console.log("margin :", margin);
    console.log("sliderWidth :", sliderWidth);
    console.log("sliderPosition :", sliderPosition);
}

//function to call all the functions
function init() {
    calcItemDisplayed();
    calcSliderWidth();
    calcSliderPosition();
    for (let i = 0; i < itemDisplayed; i++) { //for beeing sure that the slider is at the beginning
        next();
    }
}

//call the init function when the page is loaded
window.addEventListener('load', init);

//call the init function when the window is resized
window.addEventListener('resize', init);