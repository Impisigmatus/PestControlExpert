<template>
  <pce-top-bar />
  <router-view />
  <div v-show="timer" class="callnowbutton">
    <a class="phone" href="tel:+7999999999"> Связаться с нами </a>
  </div>
  <pce-footer />
  <a v-show="arrow" href="#about" class="pageup">
    <svg
      class="up"
      viewBox="0 0 26 26"
      fill="#fff"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path
        fill-rule="evenodd"
        clip-rule="evenodd"
        d="M26 0H0V26H26V0ZM4.6593 17.7519L13.1233 10.33L21.5873 17.7519L22.9059 16.2481L13.7826 8.24813L13.1233 7.67L12.464 8.24813L3.3407 16.2481L4.6593 17.7519Z"
        fill="black"
      />
    </svg>
  </a>
</template>

<script>
export default {
  data() {
    return {
      timer: false,
      arrow: false,
    };
  },
  methods: {
    changeVisibleCallButton() {
      setTimeout(() => {
        this.timer = true;
      }, 6000);
    },
    changeVisibleArrowUp() {
      window.addEventListener('scroll', () => {
        if (document.documentElement.scrollTop > 1200) {
          this.arrow = true;
        } else {
          this.arrow = false;
        }
      });
    },
  },
  mounted() {
    this.changeVisibleCallButton();
    this.changeVisibleArrowUp();
  },
};
</script>

<style>
#app {
  font-family: 'Montserrat', sans-serif;
  font-weight: 400;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

* {
  margin: 0 auto;
  transition: all 0.25s ease-out;
  box-sizing: border-box;
}
.pageup {
  width: 26px;
  height: 26px;
  position: fixed;
  bottom: 100px;
  right: 5%;
  z-index: 30;
  background-color: #fff;
  cursor: pointer;
  border: 1px solid white;
  transition: 1s all;
}
/* TODO: могут понадобится стили */
/* $main-color: #377dff; */
/* body {
  height: 100%;
} */
/* html::before {
  content: '';
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: -1;
  background: linear-gradient(-45deg, #0011ff, #00aeff, #000c75, #009ece);
  background-size: 100% 100%;
  animation: gradient 30s ease infinite;
} */
.callnowbutton {
  transition: all 0.5s all 0.5s;
  transform: rotate(90deg);
  position: fixed;
  z-index: 100;
  padding: 10px;
  border-radius: 20px 20px 0 0;
  left: -45px;
  bottom: 280px;
  height: 30px;
  text-align: center;
  font-weight: 700;
  font-size: 10px;
  overflow: hidden;
  background: #00bb00;
  display: flex;
  justify-content: center;
  align-items: center;
}
.phone {
  text-decoration: none;
  color: #fff;
}
@keyframes gradient {
  0% {
    background-position: 0% 50%;
  }
  50% {
    background-position: 100% 50%;
  }
  100% {
    background-position: 0% 50%;
  }
}
</style>
