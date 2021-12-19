<template>
  <div id='show'>
    <div v-if="show">
      <div v-for="article in articles" id='wrap_box' :key="article.id">
        <a-card :title="article.title" class='article' size="small" >
          <a slot="extra" href="javascript:;" @click="showDetail(article)"
            >查看全部内容</a
          >
          <div class='tips'>
            <div class="time">
              创建时间:{{ article.created_on}}
            </div>
            <!-- <div class="categories">
              类别：
              <a-tag v-for="t in article.categories" :key="t">{{t}}</a-tag>
            </div> -->
            <div class="tags">
              标签:
               <a-tag >{{article.tag.name}} </a-tag>
            </div>
          </div>
          <div class="context">
            <vue-markdown class='markdown-body'>{{article.content}}</vue-markdown>
          </div>
        </a-card>
        <div class="kong"></div>
      </div>
    </div>
    <div v-else>
        <Detail :article='currentArticle' @show="$data.show = true; $data.currentArticle = null " />
    </div>
  </div>
</template>

<script>
// const md = require("markdown-it")().use(require("markdown-it-mark"));
import Detail from "@/articleDetail";
import VueMarkdown from 'vue-markdown'
export default {
  name: "Show",
  components: {
    Detail,
    VueMarkdown
  },
  data() {
    return {
      articles: [],
      show: true,
      currentArticle:null
    };
  },
  async mounted() {
    await this.$api.Article.get_index().then((r) => {
      // console.log('index',r)
      this.articles = r.data.lists;
    });
  },
  methods: {
    showDetail(article){
        this.currentArticle = article 
        this.show = false
    }
  },
};
</script>

<style lang='less' scoped>
#show{
  max-width: 600px;
  margin-top: 30px;
  // #wrap_box{
  //   display: flex;
  //   justify-content: center;
  // }
}

.article{
  padding: 10px;
  border-radius: 8px;
  // background:#F2F5FC;
  box-shadow: 0 0 1px #ccc;
}
.kong {
  height: 20px;
  flex:0;
}
.context {
  height: 60px;
  overflow: hidden;
  text-overflow: ellipsis;
  padding-top: 13px;
  font-size: medium;
}
.tips{
  display: flex;
  justify-content: space-between;
  border-bottom: 1px solid #ccc;
  padding-bottom: 10px;
  font-size: smaller;
}
@media screen and (max-width:992px) {
  #show{
    margin: 0 auto;
  }
}
</style>