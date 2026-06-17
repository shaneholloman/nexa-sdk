package com.geniex.sdk.bean

/** Mirrors `geniex_ModelQueryOutput`. Result of a plan-only query. */
data class ModelQuery(
    val model_name: String,
    val runtime_id: String,
    val model_type: ModelType,
    val candidates: Array<PrecisionCandidate>,
)
