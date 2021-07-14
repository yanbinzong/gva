<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
             <el-form-item label="问题编号:"><el-input v-model.number="formData.question_id" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="问题标题:">
                <el-input v-model="formData.question_name" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="说明:">
                <el-input v-model="formData.direction" clearable placeholder="请输入" ></el-input>
          </el-form-item>


        <el-form-item label="是否必答:">
<!--          <el-switch v-model.number="formData.answer" ></el-switch>-->

          <el-switch v-model.number="formData.answer"
                     :active-value="1"
                     :inactive-value="0"
                     @change=chang($event,formData.answer)>
          </el-switch>

        </el-form-item>

             <el-form-item label="题型编号:">
                <el-input v-model="formData.topic_id" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="发题人:">
                <el-input v-model="formData.author" clearable placeholder="请输入" ></el-input>
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
    createQuestion,
    updateQuestion,
    findQuestion
} from "@/api/question";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "Question",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
            question_id:0,
            question_name:"",
            direction:"",
            answer:0,
            topic_id:"",
            author:"",
            
      }
    };
  },
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createQuestion(this.formData);
          break;
        case "update":
          res = await updateQuestion(this.formData);
          break;
        default:
          res = await createQuestion(this.formData);
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
    const res = await findQuestion({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.requestion
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