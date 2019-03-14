/**
 * @file    trans.c
 * @copyright defined in aergo/LICENSE.txt
 */

#include "common.h"

#include "ast_blk.h"
#include "trans_id.h"

#include "trans.h"

static void
trans_init(trans_t *trans, flag_t flag, ir_t *ir)
{
    trans->flag = flag;
    trans->ir = ir;

    trans->blk = NULL;

    trans->md = NULL;
    trans->fn = NULL;
    trans->bb = NULL;

    trans->cont_bb = NULL;
    trans->break_bb = NULL;

    trans->is_global = false;
}

void
trans(ast_t *ast, flag_t flag, ir_t *ir)
{
    int i;
    trans_t trans;
    ast_blk_t *root = ast->root;

    if (has_error())
        return;

    trans_init(&trans, flag, ir);

    vector_foreach(&root->ids, i) {
        id_trans(&trans, vector_get_id(&root->ids, i));
    }
}

/* end of trans.c */
