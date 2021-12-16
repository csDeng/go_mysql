<template>
    <div>
        <div id='forbidden'>
            <img  src='../../assets/404.png'/>
        </div>
        
    </div>
</template>

<script>
export default {
    name:'NotFound',
    data(){
        return {

        }
    },
    methods: {

    }
}

</script>


<style lang='less' scoped>
#forbidden{
    position: fixed;
    width: 100vw;
    height:100vh;
    display: flex;
    justify-content: center;
    align-items: center;
}

#forbidden > img {
    display: block;
    width: 500px;
    height:500px;
}
</style>