import { getDict } from "@/utils/dictionary";
export default {
    data() {
        return {
            page: 1,
            total: 10,
            pageSize: 10,
            tableData: [],
            searchInfo: {}
        }
    },
    methods: {
        filterDict(value, type) {
            const rowLabel = this[type + "Options"] && this[type + "Options"].filter(item => item.value == value)
            return rowLabel && rowLabel[0] && rowLabel[0].label
        },
        async getDict(type) {
            const dicts = await getDict(type)
            this[type + "Options"] = dicts
            return dicts
        },
        handleSizeChange(val) {
            this.pageSize = val
            this.getTableData()
        },
        handleSizeChangeQuestion(val) {
            this.pageSize = val
            this.getTableDataSelect()
        },
        handleCurrentChange(val) {
            this.page = val
            this.getTableData()
            this.getTableDataQuerstionType();
        },
        handleCurrentChangeQestion(val) {
            this.page = val
            this.getTableDataSelect()
            this.getTableDataQuerstionType();
        },
        async getTableData(page = this.page, pageSize = this.pageSize) {
            const table = await this.listApi({ page, pageSize, ...this.searchInfo })
            if (table.code == 0) {
                this.tableData = table.data.list
                this.total = table.data.total
                this.page = table.data.page
                this.pageSize = table.data.pageSize
            }
        },
        async getTableDataQuerstionType(page = this.page, pageSize = this.pageSize) {
            const table = await this.listApiQuestionType({ page, pageSize, ...this.searchInfo })
            if (table.code == 0) {
                this.tableData = table.data.list
                this.total = table.data.total
                this.page = table.data.page
                this.pageSize = table.data.pageSize
            }
        },
        async getTableDataBranch_officeList(page = this.page, pageSize = this.pageSize) {
            const table = await this.listApiBranch_officeList({ page, pageSize, ...this.searchInfo })
            if (table.code == 0) {
                this.tableData = table.data.list
                this.total = table.data.total
                this.page = table.data.page
                this.pageSize = table.data.pageSize
            }
        },
        async getTableDataSelect(page = this.page, pageSize = this.pageSize,exam_paper_id =this.formData.exam_paper_id) {
            const table = await this.listApi({ page, pageSize,exam_paper_id, ...this.searchInfo })
            if (table.code == 0) {
                this.tableData = table.data.list
                this.total = table.data.total
                this.page = table.data.page
                this.pageSize = table.data.pageSize
            }
        },

    }
}