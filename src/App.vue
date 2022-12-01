<template>
  <pce-top-bar />
  <router-view />
  <div v-show="timer" class="callnowbutton">
    <a class="phone" href="tel:+7999999999"> Связаться с нами </a>
  </div>
  <pce-footer />
</template>

<script>
export default {
  data() {
    return {
      timer: false,
    };
  },
  methods: {
    changeVisible() {
      setTimeout(() => {
        this.timer = true;
      }, 3000);
    },
  },
  mounted() {
    this.changeVisible();
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
