<template>
  <div>
    <a-card
      title="团队信息"
      :style="{
        height: '85vh',
      }"
    >
        <div class="box">
            <img src="https://img-blog.csdnimg.cn/d2a0fa23b2cb4dfa98913857c8fdcb31.png">
            
            <div id="desc">
                <h1>奥特曼不打小怪兽团队</h1>
                <p>系统架构:邓崇森</p>
                <p>前端开发:邓文杰</p>
                <p>后端开发:李潼</p>
                <p>模型设计:郑艺涵</p>
            </div>
        </div>
    </a-card>
  </div>
</template>


<script>
export default {
  name: "About",
};
</script>


<style lang='less' scoped>
.box{
    width:300px;
    margin:0 auto;
    text-align: center;
    #desc{
        padding-top: 50px;
    }
}
</style>