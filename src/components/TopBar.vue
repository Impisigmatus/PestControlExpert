<template>
  <div class="navbar">
    <img
      class="navbar__img"
      src="@/assets/logo.png"
      alt="#"
      @click="$router.push('/')"
    />
    <div class="logo">
      <h1>Pest Control Expert</h1>
    </div>
    <div class="navbar__btns"></div>
  </div>
</template>

<script>
export default {
  name: 'PceTopBar',
};
</script>

<style scoped>
.navbar {
  height: 100px;
  /* box-shadow: 1px 2px 4px rgb(120, 188, 252); */
  display: flex;
  padding: 15px 30px 15px 30px;
  flex-direction: column;
  align-items: center;
  padding: 0 15px;
  background: radial-gradient(
    circle at 24.1% 68.8%,
    rgb(50, 50, 50) 0%,
    rgb(0, 0, 0) 99.4%
  );
  border-bottom: 3px solid #ffffff;
}

.logo {
  margin-top: 5px;
  text-transform: uppercase;
  color: #ffffff;
  font-size: 12px;
  font-weight: 700;
}

.navbar__img {
  width: 70px;
  height: 60px;
}

@media (max-width: 768px) {
  .navbar__btns {
    display: none;
  }
}
</style>
