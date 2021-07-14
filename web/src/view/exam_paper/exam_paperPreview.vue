<template>
  <div class="view-layout">
    <div class="main" v-if="!isNotPublish">
      <div class="header">
        <h1>试卷状态已禁止！</h1>
      </div>
      <div class="content">
          <p>您所填写的问卷未发布，暂不能填写。</p>
      </div>
    </div>

    <div class="main" v-else-if="isExpired">
        <div class="header">
          <h1>问卷已过期！</h1>
        </div>
        <div class="content">
          <p>您所填写的问卷已到截止日期，暂不能填写。</p>
        </div>
      </div>

    <div class="main" v-if="(!isExpired && isNotPublish)">
      <div class="header">
        <h1>{{tableData.paper_title}}</h1>
      </div>

      <div class="content">
        <div class="intro">
          <p>{{tableData.intro}}</p>
          <p>截止日期：{{ tableData.end_at | renderFormatDate }}</p>
        </div>
      </div>
  <!--ParperQuestion 试卷组件-->
  <ParperQuestion :tableChildData ="this.tableData"  ref="mylistquestion">
     <ElRow>
       <div style=" text-align:center; margin: 15px;">
         <el-button type="success"
                    v-if="isAdmin"
                    @click="goBack">返回试卷列表
         </el-button>

         <el-button type="primary"
                    @click="submitQuestion">提交问卷
         </el-button>
       </div>
     </ElRow>
  </ParperQuestion>
    </div>
  </div>
</template>
<script>
import {
  createAnswer_question_front
} from "@/api/answer_question";  //  此处请自行替换地址
import {
  findExam_paperQuestion,
} from "@/api/exam_paper";  //  此处请自行替换地址

import ParperQuestion from './components/question.vue'
export default {
  components: {
     ParperQuestion,
  },
  props: [],
  data() {
    return {
      exam_paper:{
        topic:[]
      },
      tableData:'',
      finished: false,
      time_end:'',
    }
  },
  filters: {  // 前端时间格式2020-02-11T12:24:18.000+0000转化成正常格式
    renderFormatDate: function (time) {
      var dateee = new Date(time).toJSON();
      return new Date(+new Date(dateee) + 8 * 3600 * 1000).toISOString().replace(/T/g, ' ').replace(/\.[\d]{3}Z/, '')
    },
  },
  computed: {
    isAdmin () {
      return 1
      // return this.$store.getters.isAdmin
    },

    // true过期  false未过期
    isExpired () {
      const timestamp = parseInt(new Date().getTime()/1000);    // 当前时间戳
      console.log("isExpired:"+this.getTimestamp(this.time_end) , timestamp)
      return Number(this.getTimestamp(this.time_end) < timestamp)
    },

    // 判断是否 公开1 禁止0
    isNotPublish () {
      return this.tableData.paper_status
    },
  },
  watch: {},
  created() {
    this.tableData=this.getTableData()
  },
  mounted() {},
  // 与生命周期同级 provide inject
  privide() {
    return {
      //子组件调用的名字：对应的方法（当前页面，祖父级元素的方法）
     // submitaaa: this.submitForm()
    }
  },
  methods: {
    submitForm() {
      this.$refs['elForm'].validate(valid => {
        if (!valid) return
        // TODO 提交表单
      })
    },
    resetForm() {
      this.$refs['elForm'].resetFields()
    },
    // GET QUESTION ALL
    async getTableData() {

      var  exam_paper_id = this.$route.query.exam_paper_id
      const  table = await findExam_paperQuestion({ exam_paper_id: exam_paper_id })
      if (table.code == 0) {
        console.log(table.data.list)
        this.tableData = table.data.list
        this.time_end = table.data.list.end_at
      }
    },

    getTimestamp(time) { //把时间日期转成时间戳
      return (new Date(time)).getTime() / 1000
    },

    // 前端时间格式2020-02-11T12:24:18.000+0000转化成正常格式
    renderTime(date) {
      var dateee = new Date(date).toJSON();
      return new Date(+new Date(dateee) + 8 * 3600 * 1000).toISOString().replace(/T/g, ' ').replace(/\.[\d]{3}Z/, '')
    },
    goBack () {
      this.$router.push('/layout/exam_paperP/exam_paperTable')
    },
    // 提交问题
    submitQuestion(){
      // 我去调用子组件
      this.$refs.mylistquestion.childQustion()
      this.user_id = 1
      //const Id = ''
      const result = []
      console.log(this.tableData.Question)
      console.log("=======shang==========")
      this.tableData.Question.forEach((Question,index)=>{

        if((Question.topic_id ==="9")){ // 1 一个文本框 2 不要in

          const score = parseFloat(Question.selectContent)
          console.log("一个文本框Question",Question)
          const curQues = {
            user_id :this.user_id,
            paper_id :Question.exam_paper_id,
            question_id :Question.ID,
            answer_content :!Question.infoContent ? Question.selectContent:Question.selectContent ,
            score :score,
          }
          result.push(curQues)
          console.log("动态文本 curQues:"+JSON.stringify(curQues))
        }

        if((Question.topic_id ==="8")){ // 1 一个文本框 2.

          const score = parseFloat(Question.selectContent)
          console.log("一个文本框Question",Question)
          const curQues = {
            user_id :this.user_id,
            paper_id :Question.exam_paper_id,
            question_id :Question.ID,
            answer_content :Question.infoContent,
            score :score,
          }
          result.push(curQues)
          console.log("动态文本 curQues:"+JSON.stringify(curQues))
        }


        if((Question.topic_id ==="5")){ // 1 一个文本框 2.

          const score = parseFloat(Question.selectContent)
          console.log("一个文本框Question",Question)
          const curQues = {
            user_id :this.user_id,
            paper_id :Question.exam_paper_id,
            question_id :Question.ID,
            answer_content :Question.infoContent,
            score :score,
          }
          result.push(curQues)
          console.log("动态文本 curQues:"+JSON.stringify(curQues))
        }

        if((Question.topic_id ==="3")){ // 1 一个文本框 2.多图文本框

          const score = parseFloat(Question.selectContent)
          console.log("一个文本框Question",Question)
          const curQues = {
            user_id :this.user_id,
            paper_id :Question.exam_paper_id,
            question_id :Question.ID,
            answer_content :Question.imgurl,
            score :score,
          }
          result.push(curQues)
          console.log("一个文本框 curQues:"+JSON.stringify(curQues))
        }else if((Question.topic_id ==="1")){ // 1 一个文本框

          const score = parseFloat(Question.selectContent)
          console.log("一个文本框Question",Question)
          const curQues = {
            user_id :this.user_id,
            paper_id :Question.exam_paper_id,
            question_id :Question.ID,
            answer_content :Question.selectContent,
            score :score,
          }
          result.push(curQues)
          console.log("一个文本框 curQues:"+JSON.stringify(curQues))
        }else if(Question.topic_id ==="2"){  //2 三个文本框

          // //let arr=[]
          // let obj={
          //   img:value
          // }
          // this.tableChildData.Question.forEach((item)=>{
          //   if(item.topic_id ==='3'){
          //     //  arr.push(obj)
          //     this.idImgArr.push(obj)
          //     this.$set(item,'imgurl', JSON.stringify(this.idImgArr))
          //   }
          let obj={
            list:{
              dydhcs:Question.dydhcs,
              dyxzhcs:Question.dyxzhcs,
              sdkcs:Question.sdkcs,
            }
          }
          this.$set(Question,'data', JSON.stringify(obj))
            console.log("========Question======")

          var inputContent = new Array();
          inputContent['dydhcs'] = Question.dydhcs;
          inputContent['dyxzhcs'] = Question.dyxzhcs;
          inputContent['sdkcs'] = Question.sdkcs;

         var json_inputContent= {"dydhcs":Question.dydhcs,"dyxzhcs":Question.dyxzhcs,"sdkcs":Question.sdkcs}
          const score = parseFloat(Question.dydhcs)+ parseFloat(Question.dyxzhcs)+ parseFloat(Question.sdkcs)
          console.log("score"+score,json_inputContent)

          const curQues = {
               user_id :this.user_id,
               paper_id :Question.exam_paper_id,
               question_id :Question.ID,
               answer_content :Question.data,
               score :score,
          }

          result.push(curQues)
          console.log("三文本框：",(JSON.stringify(curQues)))
        }else if(Question.type ==="单选"){
          const curQues = {
            // user_id :user_id,
            // paper_id :paper_id,
            // question_id :question_id,
            // answer_content :answer_content,
            // score :score,

          }
          result.push(curQues)
          console.log("单选 curQues:"+curQues)
        }else if(Question.type ==="多选"){
          const curQues = {
            // user_id :user_id,
            // paper_id :paper_id,
            // question_id :question_id,
            // answer_content :answer_content,
            // score :score,
          }
          result.push(curQues)
          console.log("多选 curQues:"+curQues)
        }
        // else{
        //   const score = parseFloat(Question.selectContent)
        //   console.log("一个文本框Question",Question)
        //   const curQues = {
        //     //user_id :user_id,
        //     paper_id :Question.exam_paper_id,
        //     question_id :Question.ID,
        //     answer_content :Question.selectContent,
        //     score :score,
        //   }
        //   result.push(curQues)
        //   console.log("一个文本框default curQues:"+JSON.stringify(curQues))
        //
        // }
         console.log(Question,index)
      })
      console.log("all_data:",result)
      console.log('======begin result======')
      var str = JSON.stringify(result)
      console.log('======end result======')
      console.log(str)

      //防止重复提交
      //this.finished = true
      this.enterDialog(result)
    },
    // 創建文件
    async enterDialog(result) {
      let res = await createAnswer_question_front(result);
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"试卷提交成功！"
        })


        this.reload(); //此处刷新页面
        this.closeDialog();
      }
    },
  }
}

</script>
<style scoped>
.view-layout {
  background-color: rgb(237, 240, 248);
  min-height: 100%;
  max-height: 100%;
  height: 100%;
  width: 100%;
  padding: 20px 10px;
  overflow-y: scroll;
}

.view-layout .main {
  max-width: 800px;
  width: 100%;
  height: auto;
  margin: 0 auto;
  padding-bottom: 30px;
  background-color: #fff;
  -webkit-box-shadow: 0 2px 5px 1px rgba(124, 124, 124, .2);
  box-shadow: 0 2px 5px 1px rgba(124, 124, 124, .2);
}

.view-layout .header {
  padding: 30px 10px;
  height: auto;
  min-height: 10px;
  border-bottom: 2px dotted #eee;
}

.view-layout .header h1 {
  /*width: 500px;*/
  width: 100%;
  margin: 0 auto;
  margin-top: 5px;
  margin-bottom: -10px;
  text-align: center;
}

.view-layout .content {
  padding: 20px;
}

.view-layout .intro {
  font-size: 14px;
  text-indent: 2em;
}

.view-layout .footer {
  margin-top: 10px;
  padding-top: 10px;
  text-align: center;
  border-top: 1px dotted #eee;
}

.code-row-bg button {
  margin: 0 10px;
}

.user-info {
  width: 100%;
  padding: 30px 30px 0 30px;
}
</style>