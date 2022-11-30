<template>
  <pce-top-bar />
  <router-view />
</template>

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

body {
  background: linear-gradient(-45deg, #0011ff, #00aeff, #000c75, #009ece);
  background-size: 200% 200%;
  animation: gradient 30s ease infinite;
  height: 100%;
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
