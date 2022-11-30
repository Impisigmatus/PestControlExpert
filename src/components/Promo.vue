<template>
  <section id="up" class="promo">
    <div class="container">
      <header class="header">
        <a href="#" class="header__logo">
          <img src="@/assets/logo.png" alt="logo" />
        </a>
        <div class="header__official">
          Официальный представитель
          <span>Garmin, Polar, Suunto</span>
        </div>
        <div class="header__contacts">
          <a href="tel:+74999228974" class="header__phone"
            >+7 (499) 922-89-74</a
          >
          <button data-modal="consultation" class="button">
            заказать звонок
          </button>
        </div>
      </header>

      <div class="promo__wrapper">
        <h1 class="promo__header">Подбор пульсометра</h1>
        <h2 class="promo__subheader">с учётом Вашего уровня подготовки</h2>

        <div class="promo__icons">
          <div class="promo__icon">
            <img src="icons/main_screen/timer.png" alt="timer" />
            Время подбора: 10 минут
          </div>
          <div class="promo__icon">
            <img src="icons/main_screen/hand-cart.png" alt="hand-cart" />
            Бесплатная доставка
          </div>
          <div class="promo__icon">
            <img src="icons/main_screen/messages.png" alt="messages" />
            Ответим на все вопросы
          </div>
        </div>

        <button data-modal="consultation" class="button button_main">
          заказать кОНСУЛЬТАЦИЮ
        </button>
        <div class="promo__link">
          или
          <a href="#">ПЕРЕЙТИ В КАТАЛОГ</a>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
export default {
  name: 'PcePromo',
};
</script>

<style scoped>
.promo {
  min-height: 650px;
  padding: 21px 0 93px 0;
  background: url('https://xn--80aaldbt0aefothn2d5exb.xn--p1ai/wp-content/uploads/2018/09/dezinf2-5.jpg')
    center center/cover no-repeat;
}
.promo__wrapper {
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  margin-top: 90px;
}
.promo__header {
  color: #fff;
  font-size: 42px;
  font-weight: 700;
  text-transform: uppercase;
  line-height: 42px;
}
.promo__subheader {
  margin-top: 17px;
  color: #fff;
  font-size: 24px;
  line-height: 24px;
  font-weight: 500;
  text-transform: uppercase;
}
.promo__icons {
  margin-top: 65px;
  width: 425px;
  display: flex;
  justify-content: space-between;
}
.promo__icon {
  width: 108px;
  text-align: center;
  color: #fff;
  font-size: 14px;
  font-weight: 500;
}
.promo__icon img {
  display: block;
  margin: 0 auto;
  margin-bottom: 20px;
}
.promo__link {
  color: #fff;
  font-size: 16px;
  font-weight: 500;
  margin-top: 18px;
}
.promo__link a {
  color: #fff;
  text-transform: uppercase;
  text-decoration: underline;
}
</style>
