<template>
  <footer class="footer">
    <div class="footer__info">
      <!-- TODO: могут понадобится кнопки в футере
      <div class="footer__btns">
        <pce-button class="btns__footer" @click="$router.push('/price')"
          >Цены</pce-button
        >
        <pce-button
          class="btns__footer"
          style="margin-left: 20px"
          @click="$router.push('/map')"
          >Карта</pce-button
        >
      </div> -->
      <h2 class="footer__title">Юридический адрес</h2>
      <address class="footer__adr">
        Домодедово, мкр. Барыбино, ул. Агрохимиков, дом №7
      </address>
      <h2 class="footer__title_second">Контакты</h2>
      <h4 class="footer__subtitle">Телефон</h4>
      <a class="footer__phone" href="tel:+79057247681">+7 (905) 724-76-81</a>
      <h4 class="footer__subtitle">E-mail</h4>
      <a class="footer__phone" href="email:pestcontrolexpert@yandex.ru"
        >pestcontrolexpert@yandex.ru</a
      >
      <div class="footer__social">
        <a
          href="https://api.whatsapp.com/send?phone=79057247681&text=Здравствуйте!%20"
          class="footer__social-item"
        >
          <img src="@/assets/icons/whatsapp.svg" alt="whatsapp"
        /></a>
        <a
          href="mailto:pestcontrolexpert@yandex.ru?subject=PestControlExpert"
          class="footer__social-item"
        >
          <img src="@/assets/icons/email.svg" alt="email" />
        </a>
      </div>
    </div>
  </footer>
</template>

<script>
export default {
  name: 'PceFooter',
};
</script>

<style scoped>
.footer {
  position: relative;
  border-top: 5px solid #b1c8f3;
}
.footer__info {
  position: absolute;
  padding: 20px 40px 20px 40px;
  top: 0;
  left: 50%;
  transform: translateX(-50%);
  width: 100%;
  height: 300px;
  background-color: #8bc6ec;
  background-image: linear-gradient(135deg, #8bc6ec 0%, #9599e2 100%);
}
.footer__btns {
  display: flex;
}

.footer__adr {
  margin: 15px;
  font-size: 15px;
  line-height: 18px;
  font-weight: 400;
  text-align: center;
  color: #ffffff;
}
.footer__phone {
  display: block;
  color: #ffffff;
  font-size: 15px;
  font-weight: 400;
  margin-bottom: 18px;
  text-align: center;
}
.footer__phone:hover {
  color: #ffffff;
  text-decoration: none;
}

.footer__title {
  color: #ffffff;
  margin-top: 10px;
}
.footer__title_second {
  color: #ffffff;
  margin-bottom: 10px;
}
.footer__subtitle {
  color: #ffffff;
}
.footer__social-item {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 50px;
  height: 50px;
  border: 2px solid #8bc6ec;
  border-radius: 100%;
  margin: 0 15px;
  transition: 0.5s all;
}
.footer__social-item:hover {
  box-shadow: 3px 3px 20px black;
  border: none;
}
.footer__social-item img {
  height: 45px;
}
.footer__social {
  display: flex;
  justify-content: center;
  margin-bottom: 10px;
}
</style>
