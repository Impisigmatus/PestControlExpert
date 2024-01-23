<template>
  <section class="price" id="prices">
    <h2 class="title__section-title">Прайс-лист (цены указаны в рублях)</h2>
    <h3 class="title__section-title">Таблица цен на услуги до 150 кв. м</h3>
    <div class="divider"></div>
    <div class="price__wrapper">
      <div class="price__item">
        <div class="price__type">Мониторинг</div>
        <div class="price__count"></div>
        <div class="price__descr">
          <ul>
            <li>750 р.</li>
            <li>1500 р. (Премиум)</li>
          </ul>
        </div>
      </div>
      <div class="price__item">
        <div class="price__type">Дератизация</div>
        <div class="price__count"></div>
        <div class="price__descr">
          <ul>
            <li>1850 р.</li>
            <li>3000 р. (Премиум)</li>
          </ul>
        </div>
      </div>
      <div class="price__item">
        <div class="price__type">Дезинсекция</div>
        <div class="price__count"></div>
        <div class="price__descr">
          <ul>
            <li>2300 р.</li>
            <li>3900 р. (Премиум)</li>
          </ul>
        </div>
      </div>
      <div class="price__item">
        <div class="price__type">Дезинфекция</div>
        <div class="price__count"></div>
        <div class="price__descr">
          <ul>
            <li>4800 р.</li>
            <li>9900 р. (Премиум)</li>
          </ul>
        </div>
      </div>
      <div class="price__item">
        <div class="price__type">СЭО</div>
        <div class="price__count"></div>
        <div class="price__descr">
          <ul>
            <li>4000 р.</li>
            <li>8000 р. (Премиум)</li>
          </ul>
        </div>
      </div>
    </div>
    <h3 class="title__section-title">
      Цена услуги зависит: от площади, комплексности услуг, заселенности
      помещений и вида членистоногих.
    </h3>
  </section>
</template>

<script>
export default {
  name: 'PcePrice',
};
</script>

<style scoped>
.price {
  padding: 30px;
}
.price__wrapper {
  margin-top: 60px;
  padding: 0 60px;
  display: grid;
  grid-template: repeat(1, minmax(45px, auto)) / repeat(5, 265px);
  grid-auto-rows: minmax(45px, auto);
  row-gap: 15px;
  justify-content: space-between;
}
.price__item {
  position: relative;
  display: grid;
  grid-template: repeat(3, minmax(21px, auto)) / repeat(2, auto);
  row-gap: 15px;
  padding-left: 25px;
  column-gap: 5px;
}
.price__item:before {
  content: '';
  width: 15px;
  height: 15px;
  background-color: blue;
  border-radius: 100%;
  position: absolute;
  top: 0px;
}
.price__type,
.price__count {
  font-weight: bold;
  font-size: 14px;
  line-height: 21px;
}
.price__count {
  justify-self: end;
}
.price__descr {
  grid-column: 0.3333333333;
  font-size: 20px;
  line-height: 22px;
}
@media (max-width: 992px) {
  .price__wrapper {
    grid-template-columns: repeat(auto-fit, 350px);
    padding: 0;
  }
}
@media (max-width: 576px) {
  .price__wrapper {
    grid-template-columns: repeat(auto-fit, 300px);
  }
}
</style>
