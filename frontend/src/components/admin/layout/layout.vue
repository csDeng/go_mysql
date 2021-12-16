<template>
  <div id="app">
    <a-layout id="components-layout-demo-side" :style="{
      height:'100vh',

    }" >
      <my-slider />
      <my-header />
      <my-content />
    </a-layout>
  </div>
</template>

<script>
import MySlider from './MySlider.vue'
import MyHeader from './MyHeader.vue'
import MyContent from './MyContent.vue'
export default {
  name: "Layout",
  components:{
    MySlider,
    MyHeader,
    MyContent
  },
};
</script>



<style lang='less' scoped >
</style>
