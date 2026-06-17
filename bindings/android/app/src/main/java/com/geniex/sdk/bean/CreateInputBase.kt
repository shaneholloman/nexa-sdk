package com.geniex.sdk.bean

interface CreateInputBase : InputPluginBase {
    val model_path: String?
    val config: ModelConfig
    /**
     * [RuntimeIdValue] to use for the model
     */
    override val runtime_id: String?
    /**
     * Compute unit alias. `null` selects the runtime default ([ComputeUnitValue.HYBRID]
     * for `llama_cpp`, [ComputeUnitValue.NPU] for `qairt`).
     */
    override val compute_unit: String?
}