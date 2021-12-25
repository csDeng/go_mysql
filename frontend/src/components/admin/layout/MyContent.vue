<template>
    <div>
        <a-layout-content
          :style="{
            margin: '84px 20px',
            height: '85vh',
            overflow:'auto',
            minWidth:'calc( 100vw - 240px)'
          }"
        >
          <router-view />
          <Footer />
        </a-layout-content>
        
    </div>
</template>

<script>
import Footer from "../../Footer.vue"

export default {
    name:'MyContent',
    components:{
      Footer,
    }
}
</script>