<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
             <el-form-item label="用户编号:"><el-input v-model.number="formData.user_id" clearable placeholder="请输入"></el-input>
          </el-form-item>

             <el-form-item label="试卷编号:"><el-input v-model.number="formData.paper_id" clearable placeholder="请输入"></el-input>
          </el-form-item>

             <el-form-item label="问题编号:"><el-input v-model.number="formData.question_id" clearable placeholder="请输入"></el-input>
          </el-form-item>

             <el-form-item label="回答内容json:">
                <el-input v-model="formData.answer_content" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           <el-form-item>
           <el-button @click="save" type="primary">保存</el-button>
           <el-button @click="back" type="primary">返回</el-button>
           </el-form-item>
    </el-form>
</div>
</template>

<script>
import {
    createAnswer_question,
    updateAnswer_question,
    findAnswer_question
} from "@/api/answer_question";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "Answer_question",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
            user_id:0,
            paper_id:0,
            question_id:0,
            answer_content:"",
            
      }
    };
  },
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createAnswer_question(this.formData);
          break;
        case "update":
          res = await updateAnswer_question(this.formData);
          break;
        default:
          res = await createAnswer_question(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"创建/更改成功"
        })
      }
    },
    back(){
        this.$router.go(-1)
    }
  },
  async created() {
   // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if(this.$route.query.id){
    const res = await findAnswer_question({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.reanswer_question
       this.type = "update"
     }
    }else{
     this.type = "create"
   }
  
}
};
</script>

<style>
</style>