<template>
  <div class="callnowbutton">
    <a href="tel:+78007751647" id="callnowbutton">
      <img alt="Call Now Button" width="40" />
      <span style="color: #ffffff">Бесплатная консультация</span>
    </a>
  </div>
</template>

<script>
export default {
  name: 'PceCons',
};
</script>

<style scoped>
.callnowbutton {
  display: block;
  position: fixed;
  text-decoration: none;
  z-index: 2147483647;
  width: 100%;
  left: 0;
  bottom: 0;
  height: 60px;
  text-align: center;
  color: #fff;
  font-weight: 600;
  font-size: 120%;
  overflow: hidden;
  background: #00bb00;
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>
