<template>
  <div id="list">
    <a-list
      class="demo-loadmore-list"
      :loading="loading"
      item-layout="horizontal"
      :data-source="data"
      v-if="!detail"
    >
      <div
        v-if="showLoadingMore"
        slot="loadMore"
        :style="{
          textAlign: 'center',
          marginTop: '12px',
          height: '32px',
          lineHeight: '32px',
        }"
      >
        <a-spin v-if="loadingMore" />
        <a-button v-else @click="onLoadMore"> loading more </a-button>
      </div>
      <a-list-item slot="renderItem" slot-scope="item">
        <a slot="actions" @click="show(item)">查看</a>
        <a-list-item-meta :description="item.content.slice(0, 300)">
          <a slot="title" href="https://www.antdv.com/">{{ item.title }}</a>
          <a-avatar slot="avatar" icon="book" />
        </a-list-item-meta>
      </a-list-item>
    </a-list>
    <article-detail
      v-else
      :article="currentArticle"
      @show="
        $data.detail = false;
        $data.currentArticle = {};
      "
    />
  </div>
</template>

<script>
import ArticleDetail from "../articleDetail.vue";
export default {
  name: "ArticleList",
  components: {
    ArticleDetail,
  },
  data() {
    return {
      loading: true,
      loadingMore: false,
      showLoadingMore: true,
      data: [],
      detail: false,
      currentArticle: {},
      pid: 1,
    };
  },
  mounted() {
    this.getData((res) => {
      // console.log("getdata=>", res);
      this.loading = false;
      this.data = res;
    });
  },
  methods: {
    show(article) {
      console.log(article);
      (this.currentArticle = article), (this.detail = true);
    },
    getData(callback, next = false) {
      if (next) {
        this.pid = this.pid + 1;
      }
      const { pid } = this;
      this.$api.Article.get_list(pid).then((r) => {
        if (r.data.length === 0) {
          this.$message.info("没有更多咯");
        }
        callback(r.data.lists);
      });
    },
    onLoadMore() {
      this.loadingMore = true;
      this.getData((res) => {
        if (res.length === 0) {
          this.$message.warning("没有更多咯");
        } else {
          this.data = this.data.concat(res);
          this.$nextTick(() => {
            window.dispatchEvent(new Event("resize"));
          });
        }
        this.loadingMore = false;
      }, true);
    },
  },
};
</script>
<style scoped>
#list {
  margin: 30px auto;
  max-width: 985px;
}
.demo-loadmore-list {
  min-height: 350px;
}
</style>
