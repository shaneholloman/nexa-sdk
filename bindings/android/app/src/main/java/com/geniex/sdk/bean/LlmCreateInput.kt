package com.geniex.sdk.bean

data class LlmCreateInput(
    /**
     * Optional model identifier. The QAIRT plugin no longer needs it — it
     * dispatches by reading `metadata.json` from the bundle directory. Pass
     * a value only as a hint to `geniex_resolve_device` for the gpt-oss
     * llama_cpp device-default override, or for diagnostics.
     */
    val model_name: String? = null,
    override val model_path: String,
    val tokenizer_path: String? = null,
    override val config: ModelConfig,
    /**
     * [RuntimeIdValue] to use for the model
     */
    override val runtime_id: String? = null,
    /**
     * Compute unit alias. `null` selects the runtime default ([ComputeUnitValue.HYBRID]
     * for `llama_cpp`, [ComputeUnitValue.NPU] for `qairt`). Use
     * [ComputeUnitValue.CPU] / [ComputeUnitValue.GPU] / [ComputeUnitValue.NPU] /
     * [ComputeUnitValue.HYBRID] to pin a specific compute unit.
     */
    override val compute_unit: String? = null,
) : CreateInputBase
