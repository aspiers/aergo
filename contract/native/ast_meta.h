/**
 * @file    ast_meta.h
 * @copyright defined in aergo/LICENSE.txt
 */

#ifndef _AST_META_H
#define _AST_META_H

#include "common.h"

#ifndef _AST_META_T
#define _AST_META_T
typedef struct ast_meta_s ast_meta_t;
#endif /* ! _AST_META_T */

typedef enum type_e {
    TYPE_NONE       = 0,
    TYPE_ACCOUNT,
    TYPE_BOOL,
    TYPE_BYTE,
    TYPE_FLOAT,
    TYPE_DOUBLE,
    TYPE_INT8,
    TYPE_INT16,
    TYPE_INT32,
    TYPE_INT64,
    TYPE_UINT8,
    TYPE_UINT16,
    TYPE_UINT32,
    TYPE_UINT64,
    TYPE_STRING,
    TYPE_MAP,
    TYPE_STRUCT,
    TYPE_MAX
} type_t; 

typedef enum scope_e {
    SCOPE_GLOBAL    = 0,
    SCOPE_LOCAL,
    SCOPE_SHARED,
    SCOPE_MAX
} scope_t;

typedef struct meta_struct_s {
    char *name;
} meta_struct_t;

typedef struct meta_map_s {
    type_t k_type;
    type_t v_type;
} meta_map_t;

struct ast_meta_s {
    type_t type;
    scope_t scope;
    bool is_const;
    int arr_size;

    union {
        meta_struct_t u_st;
        meta_map_t u_map;
    };
};

static inline void
ast_meta_init(ast_meta_t *meta)
{
    memset(meta, 0x00, sizeof(ast_meta_t));
}

#endif /* ! _AST_META_H */