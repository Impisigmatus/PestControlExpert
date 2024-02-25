<template>
  <footer class="footer">
    <div class="footer__info">
      <div class="footer__logo">
        <img
          class="footer__logo_img"
          src="@/assets/logo.png"
          alt="#"
          @click="$router.push('/')"
        />
        <div class="footer__logo_title">Pest Control Expert</div>
      </div>
      <div class="footer__links">
        <a href="#about" class="footer__links_page">О нас</a>
        <a href="#services" class="footer__links_page">Наши услуги</a>
        <a href="#prices" class="footer__links_page">Цены</a>
        <a href="#form" class="footer__links_page">Форма обратной связи</a>
        <a href="#map" class="footer__links_page">Карта покрытия</a>
        <a href="./privacy.html" class="footer__links_privacy"
          >Политика конфиденциальности
        </a>
      </div>
      <div class="footer__divider"></div>
      <div class="footer__contacts">
        <h4 class="footer__title">Юридический адрес</h4>
        <address class="footer__adr">
          Домодедово, мкр. Барыбино, ул. Агрохимиков, дом №7
        </address>
        <h4 class="footer__subtitle">Телефон</h4>
        <a class="footer__phone" href="tel:+79057247681">+7 (905) 724-76-81</a>
        <h4 class="footer__subtitle">E-mail</h4>
        <a class="footer__phone" href="email:pestcontrolexpert@yandex.ru"
          >pestcontrolexpert@yandex.ru
        </a>
        <div class="footer__social">
          <a href="https://t.me/pestcontrolexpert" class="footer__social-item">
            <img src="@/assets/icons/telegram.png" alt="telegram"
          /></a>
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
    </div>
  </footer>
</template>

<script>
export default {
  name: 'PceFooter',
};
</script>

<style>
.footer {
  border-top: 3px solid #ffffff;
}
a {
  text-decoration: none;
}

.footer__logo {
  display: flex;
  flex-direction: column;
  justify-content: center;
}
.footer__divider {
  border-left: 1px solid white;
}
.footer__links {
  display: flex;
  flex-direction: column;
  justify-content: space-around;
  align-items: flex-start;
}
.footer__links_privacy {
  /* border-top: 1px solid white; */
  color: #ffffff;
}
.footer__links_page {
  display: flex;
  color: #ffffff;
}
.footer__contacts {
  display: flex;
  flex-direction: column;
  gap: 10px;
}
.footer__info {
  padding: 20px 40px 20px 40px;
  align-content: space-around;
  display: flex;
  justify-content: flex-end;
  flex-wrap: wrap;
  flex-direction: row;
  width: 100%;
  height: 300px;
  background: radial-gradient(
    circle at 24.1% 68.8%,
    rgb(50, 50, 50) 0%,
    rgb(0, 0, 0) 99.4%
  );
}
.footer__btns {
  display: flex;
}
.footer__logo_img {
  width: 128px;
  height: 128px;
}
.footer__adr {
  display: flex;
  justify-content: flex-end;
  font-size: 15px;
  line-height: 18px;
  font-weight: 400;
  color: #ffffff;
}
.footer__logo_title {
  margin-top: 5px;
  text-transform: uppercase;
  color: #ffffff;
  font-size: 12px;
  font-weight: 700;
}
.footer__phone {
  color: #ffffff;
  font-size: 15px;
  font-weight: 400;
}
.footer__phone:hover {
  color: #ffffff;
  text-decoration: none;
}

.footer__title {
  color: #ffffff;
}
.footer__title_second {
  color: #ffffff;
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
}

@media (max-width: 768px) {
  .footer__divider {
    display: none;
  }
  .footer__links {
    display: none;
  }
  .footer__logo_title {
    display: none;
  }
  .footer__info {
    height: 320px;
  }
}
</style>
