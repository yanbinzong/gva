<template xmlns="http://www.w3.org/1999/html" xmlns="http://www.w3.org/1999/html">
  <ElRow class="question-wrapper inputDeep">

    <el-col :span="24" class="question-list">
      <p v-if="tableChildData.length == 0" style="margin: 10px 0">一点东西都没有，赶快点击上方按钮添加题目吧！</p>

      <ElRow type="flex" justify="start" align="top" v-for="(question, index) in tableChildData.Question" class="question-item" :key="question.ID">
        <el-col :span="6" style="width: 60px; text-align: center">
              <h2>Q{{ index + 1 }}:</h2>
              <div class="question-action" v-show="isPreview" @click="delQuestion(index)">
              </div>
        </el-col>

        <el-col :span="18">
          <h3>[{{ question.topic_id }}] {{ question.question_name}} {{question.answer ? "（必填）" : "（选填）"}} <em style="color: #f00;" v-if="question.answer">*</em></h3>
          <p class="question-desc" v-if="question.direction !== ''">说明：{{ question.direction }}</p>

          <div class="question-options">
            <!--三文本框-->
            <div v-if="question.topic_id ==='2'" >
              <el-input v-model="question.dydhcs" placeholder="党员大会次数"  oninput = "value=value.replace(/[^\d]/g,'')"></el-input>
              <el-input v-model="question.dyxzhcs" placeholder="党员小组会次数"  oninput = "value=value.replace(/[^\d]/g,'')"></el-input>
              <el-input v-model="question.sdkcs" placeholder="上党课次数"  oninput = "value=value.replace(/[^\d]/g,'')"></el-input>

<!--              <div class="option-item" v-for="(option, index) in addCheckbox_form.options">-->
<!--                <el-row>-->
<!--                  <ElCol span="18">-->
<!--                  <el-input v-model="question.selectContent3" placeholder="上党课次数" ></el-input>-->
<!--                  <ElCol span="6">-->
<!--                    <div class="option-btn">-->
<!--                      <el-button type="success" icon="plus-round" size="small"-->
<!--                              @click="addOption(addCheckbox_form.options)"></el-button>-->
<!--                      <el-button type="warning" icon="close" size="small"></el-button>-->
<!--                    </div>-->
<!--                  </ElCol>-->
<!--                </el-row>-->
<!--              </div>-->

            </div>

            <!--判断是否有单图片-->
            <div v-if="question.topic_id ==='3'">
<!--                  原始数据{{question}}-->
<!--                  </br>-->
<!--                        {{idImgArr}}-->


            <el-input v-model="question.selectContent" placeholder="请输入..."  oninput = "value=value.replace(/[^\d]/g,'')"></el-input>
            <Upload @extend_data="getChildData"  typeString="本单位级">1</Upload>
            </div>

            <!--三个文本框加时间-->
            <div v-if="question.topic_id ==='5'">
            <el-input v-model="question.selectContent" placeholder="请输入..." oninput = "value=value.replace(/[^\d]/g,'')" ></el-input>
              <Textmore></Textmore>
            </div>
             <!--判断多选-->

            <!--判断文本-->
            <div v-if="question.topic_id ==='1'">
                <el-input v-model="question.selectContent" placeholder="请输入..." oninput = "value=value.replace(/[^\d]/g,'')"></el-input>
            </div>
            <!--判断图片-->

            <!--五图-->
            <div v-if="question.topic_id=='6'">
              <el-input v-model="question.selectContent" placeholder="请输入..."  oninput = "value=value.replace(/[^\d]/g,'')"></el-input>
              <span>本单位级</span>
              <Upload @extend_data="getChildData"  typeString="本单位级"></Upload>
              <span>公司级</span>
              <Upload @extend_data="getChildData"  typeString="公司级"></Upload>
              <span>集团级</span>
              <Upload @extend_data="getChildData"  typeString="集团级"></Upload>
              <span>市级</span>
              <Upload @extend_data="getChildData"  typeString="市级"></Upload>
              <span>国家级</span>
              <Upload @extend_data="getChildData"  typeString="国家级"></Upload>
            </div>

            <div v-if="question.topic_id=='7'">
              <span>证书图片</span>
              <Upload @extend_data="getChildData" typeString="证书图片"></Upload>
            </div>

            <!--内容加日期-->
            <div v-if="question.topic_id=='8'">
              <el-input v-model="question.selectContent" placeholder="请输入..."  oninput = "value=value.replace(/[^\d]/g,'')"></el-input>
              <Textdate></Textdate>
            </div>

            <!--下拉菜单-->
            <div v-if="question.topic_id=='9'">
              <el-input v-model="question.selectContent" placeholder="请输入..." oninput = "value=value.replace(/[^\d]/g,'')"></el-input>
            </div>

          </div>
        </el-col>
       </ElRow>

      <!--卡槽 传递按钮-->
      <div class="question-btns">
        <slot></slot>
      </div>
    </el-col>

  </ElRow>
</template>
<script>
import globalBus from '@/utils/global-bus'
import Upload from './upload.vue'
import Textmore from './textmore.vue'
import Textdate from './textdate.vue'
export default {
  components: {
    Upload,
    Textmore,
    Textdate,
  },
  props:['tableChildData'],
  data() {
    return {
      question:{
        imgurl:'',
        info:'',
        info2:'',

      },
      idImgArr:[],
      infoArr:[],
      addRadio_form: {},
      addCheckbox_form:{},
      addTextarea_form:{},
      environmentForm: {
        items: [{
          name: '',
          variable: '',
          description: ''
        }]
      },
      list:'',
      isPreview:'',
    }
  },
  computed: {},
  watch: {},
  created() {
   // console.log('2')
    this.getTextMore()
   // this.childQustion()
  },
  mounted() {},
  methods: {
    // 新建题目
    addRadio () {
      this.addRadio_modal = true
      const radioQues = {
        question: '单选题目',
        options: [],
        description: '',
        type: '单选',
        isRequired: true,
        selectContent: '',
        additional: ''
      }
      this.addRadio_form = Object.assign({}, radioQues)
      const tempData = {
        content: '选项',
        isAddition: false
      }
      this.addRadio_form.options.splice(0, this.addRadio_form.options.length, Object.assign({}, tempData))
    },
    addCheckbox () {
      this.addCheckbox_modal = true
      const checkboxQues = {
        question: '多选题目',
        options: [],
        description: '',
        type: '多选',
        isRequired: true,
        selectMultipleContent: []
      }
      this.addCheckbox_form = Object.assign({}, checkboxQues)
      const tempData = {
        content: '选项',
        isAddition: false
      }
      this.addCheckbox_form.options.splice(0, this.addCheckbox_form.options.length, Object.assign({}, tempData))
    },
    addTextarea () {
      this.addTextarea_modal = true
      const TextareaQues = {
        question: '文本题目',
        description: '',
        type: '文本',
        isRequired: true,
        selectContent: ''
      }
      this.addTextarea_form = Object.assign({}, TextareaQues)
    },
    // 新增选项
    addOption (source) {
      const tempData = {
        content: '选项',
        isAddition: false
      }
      source.push(Object.assign({}, tempData))
    },
    // 删除选项
    delOption (source, index) {
      if (source.length > 1) {
        source.splice(index, 1)
      } else {
        this.$Message.warning('最后一个啦，不要删除哦')
      }
    },
    // 提交题目
     //图片返回地址 bangmang
    // getChildData(value){
    //   this.question.imgurl = this.idImgArr;
    //   this.idImgArr.push(value)
    //   let arr=[]
    // this.tableChildData.Question.forEach((item)=>{
    //   let obj={
    //     img:value
    //   }
    //   arr.push(obj)
    // })
    //   console.log(arr)
    //   this.question.push(arr)
    //   alert(value)
    // }

    // 自己修改
    getChildData(value){

      this.question.imgurl = value

      //let arr=[]
      let obj={
        list:value
      }
      this.tableChildData.Question.forEach((item)=>{
        if(item.topic_id ==='3'){
        //  arr.push(obj)
          this.idImgArr.push(obj)
          this.$set(item,'imgurl', JSON.stringify(this.idImgArr))
        }

      })
     // console.log(arr)
     //  console.log(this.idImgArr)
     //  alert(value)
    },
    getTextMore(){

      // 插入字段
      // let obj = {
      //   info:'2222222222222222'
      // }
      // this.info.push(obj)
      // this.$set("item",'info',JSON.stringify(this.info))

      globalBus.$on('extend_textmore', (data) => {
        console.log('extend_textmore', data)
        this.question.info = data
      })

      globalBus.$on('extend_textdate', (data) => {
        console.log('extend_textdate', data)
        this.question.info2=data
      })
    },

    // 父类调用我
    childQustion(){
      globalBus.$emit('questiontext')
      // 插入字段

      this.tableChildData.Question.forEach((item)=>{
        if(item.topic_id ==='5'){
          this.$set(item,'infoContent', JSON.stringify(this.question.info))
        }
        if(item.topic_id ==='8'){
          console.log("topic8",this.question.info2)
          this.$set(item,'infoContent',JSON.stringify(this.question.info2))
        }

        if(item.topic_id ==='7'){
          this.$set(item,'infoContent', JSON.stringify(this.question.info))
        }
      })
    },
    //sse
  },
}

</script>

<style scoped>
.question-list {
  padding: 30px 0;
  overflow: visible;
}

.checkbox-list label {
  display: block;
}

.question-options {
  padding: 5px 0;
}

.option-item {
  margin: 5px 0;
}

.option-item1 {
  margin: 5px 0;
}
.question-desc {
  padding: 5px 0;
}

.question-action {
  display: none;
  color: #a9afb2;
}

.question-item {
  padding: 20px 0;
  min-width: 300px;
}

.question-item:hover {
  background: rgba(238, 238, 238, 0.47);
}

.question-item:hover .question-action {
  display: block;
  margin-top: 8px;
}

.question-action:hover {
  color: #018fe5;
  cursor: pointer;
}

.option-item {

}

.question-btns {
  margin-top: 20px;
}

.option-action {
  display: none;
  margin-left: 100px;
  padding: 0 10px;
  text-align: right;
  color: #a9afb2;
}

.option-action:hover {
  color: #018fe5;
}


.option-item:hover .option-action {
  display: inline-block;
  cursor: pointer;
}
.option-item:hover .option-action {
  display: inline-block;
  cursor: pointer;
}
.  /* 利用穿透，设置input边框隐藏 */
.inputDeep>>>.el-input__inner {
  border: 0;
}
/* 如果你的 el-input type 设置成textarea ，就要用这个了 */
.inputDeep>>>.el-textarea__inner {
  border: 0;
  resize: none;/* 这个是去掉 textarea 下面拉伸的那个标志，如下图 */
}
</style>
