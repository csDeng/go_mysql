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
        </a-layout-content>
    </div>
</template>

<script>
export default {
    name:'MyContent'
}
</script>